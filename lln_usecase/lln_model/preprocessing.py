import numpy as np
import pandas as pd
from scapy.all import *
from sklearn.preprocessing import StandardScaler
import joblib

class NetworkDataPreprocessor:
    def __init__(self, window_size=100):
        self.window_size = window_size
        self.scaler = StandardScaler()
        self.num_features = 8  # Number of features extracted from each packet
        
    def extract_features(self, packet):
        """Extract relevant features from a network packet"""
        # Initialize features with default values
        features = np.zeros(self.num_features, dtype=np.float32)
        
        # Basic packet features
        if IP in packet:
            features[0] = float(packet[IP].len)  # Packet length
            features[1] = float(packet[IP].ttl)  # Time to live
            features[2] = float(packet[IP].proto)  # Protocol
            
            if TCP in packet:
                features[3] = float(packet[TCP].sport)  # Source port
                features[4] = float(packet[TCP].dport)  # Destination port
                features[5] = float(packet[TCP].window)  # Window size
                # Convert TCP flags to numeric value
                flags = packet[TCP].flags
                features[6] = float(int(flags))  # Convert flags to integer
            elif UDP in packet:
                features[3] = float(packet[UDP].sport)  # Source port
                features[4] = float(packet[UDP].dport)  # Destination port
                features[7] = float(packet[UDP].len)  # UDP length
        
        return features
    
    def process_pcap(self, pcap_file):
        """Process a pcap file and extract features"""
        packets = rdpcap(pcap_file)
        features_list = []
        
        for packet in packets:
            features = self.extract_features(packet)
            features_list.append(features)
        
        return np.array(features_list, dtype=np.float32)
    
    def create_sequences(self, data):
        """Create sequences of fixed window size for the LNN model"""
        if len(data) < self.window_size:
            # Pad the data if it's shorter than the window size
            padding = np.zeros((self.window_size - len(data), self.num_features), dtype=np.float32)
            data = np.vstack([padding, data])
        
        sequences = []
        for i in range(len(data) - self.window_size + 1):
            sequence = data[i:i + self.window_size]
            sequences.append(sequence)
        
        sequences = np.array(sequences, dtype=np.float32)
        # Reshape to (batch_size, window_size, num_features)
        return sequences.reshape(-1, self.window_size, self.num_features)
    
    def preprocess_data(self, data):
        """Preprocess the data for model input"""
        # Ensure data is float32
        data = data.astype(np.float32)
        
        # Normalize the data
        normalized_data = self.scaler.fit_transform(data)
        
        # Create sequences
        sequences = self.create_sequences(normalized_data)
        
        return sequences
    
    def save_scaler(self, path):
        """Save the scaler for later use"""
        joblib.dump(self.scaler, path)
    
    def load_scaler(self, path):
        """Load a saved scaler"""
        self.scaler = joblib.load(path) 