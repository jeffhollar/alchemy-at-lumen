import os
import sys

# Add the parent directory to the Python path
sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

from lln_model.preprocessing import NetworkDataPreprocessor
import numpy as np
from sklearn.model_selection import train_test_split
import matplotlib.pyplot as plt
import argparse
from sklearn.ensemble import IsolationForest
import joblib


def train_model(pcap_file, validation_split=0.2, contamination=0.1):  # contamination is the proportion of outliers
	# Initialize preprocessor
	preprocessor = NetworkDataPreprocessor(window_size=100)

	# Process the pcap file
	print("Processing network traffic data...")
	features = preprocessor.process_pcap(pcap_file)

	# Create sequences
	X = preprocessor.preprocess_data(features)
	X = X.reshape(X.shape[0], -1)  # Flatten the data for IsolationForest

	# Split the data
	X_train, X_val = train_test_split(
		X, test_size=validation_split, random_state=42
	)

	# Train the Isolation Forest model
	print("Training the Isolation Forest model...")
	model = IsolationForest(contamination=contamination, random_state=42)
	model.fit(X_train)

	# Create lln_model directory if it doesn't exist
	os.makedirs('lln_model', exist_ok=True)

	# Save the model and scaler
	joblib.dump(model, 'lln_model/isolation_forest.joblib')
	preprocessor.save_scaler('lln_model/scaler.joblib')

	# Evaluate on validation set
	y_pred_val = model.predict(X_val)
	n_errors = (y_pred_val != 1).sum()  # IsolationForest labels normal as 1, anomalies as -1
	print(f"Validation set errors: {n_errors} out of {len(X_val)}")

	print("\nTraining completed!")
	print(f"Model saved to: lln_model/isolation_forest.joblib")

	return model, preprocessor


def main():
	parser = argparse.ArgumentParser(description='Train Isolation Forest model for network traffic anomaly detection')
	parser.add_argument('--pcap-file', required=True, help='Path to the PCAP file for training')
	parser.add_argument('--validation-split', type=float, default=0.2, help='Validation split ratio')
	parser.add_argument('--contamination', type=float, default=0.1, help='Expected proportion of outliers in the data')

	args = parser.parse_args()

	model, preprocessor = train_model(
		pcap_file=args.pcap_file,
		validation_split=args.validation_split,
		contamination=args.contamination
	)


if __name__ == "__main__":
	main()
