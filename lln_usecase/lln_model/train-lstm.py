import os
import sys

# Add the parent directory to the Python path
sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

import tensorflow as tf
from lln_model.model import create_model
from lln_model.preprocessing import NetworkDataPreprocessor
import numpy as np
from sklearn.model_selection import train_test_split
import matplotlib.pyplot as plt
import argparse

def train_model(pcap_file, epochs=50, batch_size=32, validation_split=0.2, learning_rate=0.001):
    # Initialize preprocessor
    preprocessor = NetworkDataPreprocessor(window_size=100)
    
    # Process the pcap file
    print("Processing network traffic data...")
    features = preprocessor.process_pcap(pcap_file)
    
    # Create sequences and labels (assuming normal traffic is labeled as 0)
    X = preprocessor.preprocess_data(features)
    y = np.zeros(len(X), dtype=np.int32)  # All normal traffic initially
    
    # Split the data
    X_train, X_val, y_train, y_val = train_test_split(
        X, y, test_size=validation_split, random_state=42
    )
    
    # Create and compile the model
    input_shape = (X_train.shape[1], X_train.shape[2])
    model = create_model(input_shape, learning_rate=learning_rate)
    
    # Train the model
    print("Training the model...")
    history = model.fit(
        X_train, y_train,
        epochs=epochs,
        batch_size=batch_size,
        validation_data=(X_val, y_val),
        callbacks=[
            tf.keras.callbacks.EarlyStopping(
                monitor='val_loss',
                patience=5,
                restore_best_weights=True
            )
        ]
    )
    
    # Create saved_model directory if it doesn't exist
    os.makedirs('lln_model/saved_model', exist_ok=True)
    
    # Save the model and scaler
    model.save('lln_model/saved_model/model.keras')
    preprocessor.save_scaler('lln_model/scaler.joblib')
    
    # Plot training history
    plt.figure(figsize=(12, 4))
    
    plt.subplot(1, 2, 1)
    plt.plot(history.history['accuracy'], label='Training Accuracy')
    plt.plot(history.history['val_accuracy'], label='Validation Accuracy')
    plt.title('Model Accuracy')
    plt.xlabel('Epoch')
    plt.ylabel('Accuracy')
    plt.legend()
    
    plt.subplot(1, 2, 2)
    plt.plot(history.history['loss'], label='Training Loss')
    plt.plot(history.history['val_loss'], label='Validation Loss')
    plt.title('Model Loss')
    plt.xlabel('Epoch')
    plt.ylabel('Loss')
    plt.legend()
    
    plt.tight_layout()
    plt.savefig('lln_model/training_history.png')
    
    print("\nTraining completed!")
    print(f"Model saved to: lln_model/saved_model")
    print(f"Training history plot saved to: lln_model/training_history.png")
    
    return model, preprocessor

def main():
    parser = argparse.ArgumentParser(description='Train LNN model for network traffic anomaly detection')
    parser.add_argument('--pcap-file', required=True, help='Path to the PCAP file for training')
    parser.add_argument('--epochs', type=int, default=50, help='Number of training epochs')
    parser.add_argument('--batch-size', type=int, default=32, help='Training batch size')
    parser.add_argument('--validation-split', type=float, default=0.2, help='Validation split ratio')
    parser.add_argument('--learning-rate', type=float, default=0.001, help='Learning rate for the optimizer')
    
    args = parser.parse_args()
    
    model, preprocessor = train_model(
        pcap_file=args.pcap_file,
        epochs=args.epochs,
        batch_size=args.batch_size,
        validation_split=args.validation_split,
        learning_rate=args.learning_rate
    )

if __name__ == "__main__":
    main() 