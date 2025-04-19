import os
import sys

# Add the parent directory to the Python path
sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

import tensorflow as tf
import numpy as np
from lln_model.preprocessing import NetworkDataPreprocessor
import joblib
import matplotlib.pyplot as plt
import argparse

class AnomalyDetector:
    def __init__(self, model_path='lln_model/saved_model/model.keras', scaler_path='lln_model/scaler.joblib'):
        try:
            self.model = tf.keras.models.load_model(model_path)
        except ValueError as e:
            print(f"Warning: Could not load model directly. Attempting to load as TensorFlow SavedModel: {e}")
            try:
                self.model = tf.keras.layers.TFSMLayer(model_path, call_endpoint='serving_default')
            except Exception as e:
                print(f"Error loading model: {e}")
                raise
                
        self.preprocessor = NetworkDataPreprocessor(window_size=100)
        self.preprocessor.load_scaler(scaler_path)
        
    def detect_anomalies(self, pcap_file, threshold=0.5):
        """Detect anomalies in network traffic"""
        # Process the pcap file
        features = self.preprocessor.process_pcap(pcap_file)
        sequences = self.preprocessor.preprocess_data(features)
        
        # Get predictions
        predictions = self.model.predict(sequences)
        
        # Calculate anomaly scores (using the probability of the normal class)
        anomaly_scores = 1 - predictions[:, 0]
        
        # Identify anomalies
        anomalies = anomaly_scores > threshold
        anomaly_indices = np.where(anomalies)[0]
        
        return {
            'anomaly_scores': anomaly_scores,
            'anomaly_indices': anomaly_indices,
            'total_anomalies': len(anomaly_indices),
            'total_packets': len(sequences)
        }
    
    def plot_anomalies(self, anomaly_scores, save_path='lln_model/anomaly_detection.png'):
        """Plot anomaly scores over time"""
        plt.figure(figsize=(12, 6))
        plt.plot(anomaly_scores, label='Anomaly Score')
        plt.axhline(y=0.5, color='r', linestyle='--', label='Threshold')
        plt.title('Network Traffic Anomaly Detection')
        plt.xlabel('Packet Sequence')
        plt.ylabel('Anomaly Score')
        plt.legend()
        plt.grid(True)
        plt.savefig(save_path)
        plt.close()

def main():
    parser = argparse.ArgumentParser(description='Detect anomalies in network traffic using trained LNN model')
    parser.add_argument('--pcap-file', required=True, help='Path to the PCAP file to analyze')
    parser.add_argument('--threshold', type=float, default=0.5, help='Anomaly detection threshold')
    parser.add_argument('--model-path', default='lln_model/saved_model/model.keras', help='Path to the trained model')
    parser.add_argument('--scaler-path', default='lln_model/scaler.joblib', help='Path to the saved scaler')
    
    args = parser.parse_args()
    
    # Initialize the detector
    detector = AnomalyDetector(model_path=args.model_path, scaler_path=args.scaler_path)
    
    # Detect anomalies
    results = detector.detect_anomalies(args.pcap_file, threshold=args.threshold)
    
    # Print results
    print(f"\nAnalysis Results:")
    print(f"Total packets analyzed: {results['total_packets']}")
    print(f"Number of anomalies detected: {results['total_anomalies']}")
    print(f"Anomaly rate: {results['total_anomalies'] / results['total_packets']:.2%}")
    
    # Plot the results
    detector.plot_anomalies(results['anomaly_scores'])
    print(f"\nAnomaly detection plot saved to: lln_model/anomaly_detection.png")

if __name__ == "__main__":
    main() 