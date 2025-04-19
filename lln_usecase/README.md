# Long-lived Network (LLN) Anomaly Detection

This project implements a deep learning model for detecting anomalies in network traffic using TensorFlow.

## Environment Setup

### Prerequisites

- macOS with Apple Silicon (M1/M2/M3)
- Homebrew package manager
- Internet connection

### Installation Steps

1. Install Miniforge (if not already installed):
```bash
brew install miniforge
```

2. Initialize conda for your shell:
```bash
conda init "$(basename "${SHELL}")"
```

3. Create a new conda environment with TensorFlow and required dependencies:

```bash
conda create -n tf2 python=3.9 tensorflow=2.9 numpy scikit-learn matplotlib joblib scapy pandas h5py -c conda-forge
```

4. Activate the conda environment:

```bash
conda activate tf2
```

## Usage

### Training the Model

To train the model on your network traffic data:

```bash
python lln_model/train.py \
    --pcap-file ./path/to/your/traffic.pcap \
    --epochs 50 \
    --batch-size 32 \
    --validation-split 0.2 \
    --learning-rate 0.001
```

Parameters:

- `--pcap-file`: Path to the PCAP file containing network traffic data
- `--epochs`: Number of training epochs (default: 50)
- `--batch-size`: Size of training batches (default: 32)
- `--validation-split`: Fraction of data to use for validation (default: 0.2)
- `--learning-rate`: Learning rate for the optimizer (default: 0.001)

### Model Output

After training, the model will save:

- Trained model at `lln_model/saved_model/model.keras`
- Data scaler at `lln_model/scaler.joblib`
- Training history plot at `lln_model/training_history.png`

## Troubleshooting

### Common Issues

1. If you encounter TensorFlow compatibility issues:
   - Make sure you're using the conda environment with TensorFlow 2.9
   - Verify you're using Python 3.9 from the conda environment

2. If you see memory-related errors:
   - Try reducing the batch size
   - Consider using a smaller subset of your data for training

### Environment Management

To manage your conda environment:

```bash
# List all conda environments
conda env list

# Activate the environment
conda activate tf2

# Deactivate the environment
conda deactivate

# Remove the environment if needed
conda env remove -n tf2
```

## Project Structure

```
.
├── README.md
├── lln_model/
│   ├── train.py           # Training script
│   ├── model.py           # Model architecture
│   ├── preprocessing.py   # Data preprocessing
│   ├── saved_model/      # Saved model directory
│   └── scaler.joblib     # Saved data scaler
```

## Features

- **Network Traffic Analysis**
  - Real-time processing of PCAP files
  - Feature extraction from network packets using Scapy
  - Support for various network protocols
  - Automated packet sequence generation

- **Advanced ML Architecture**
  - Custom LNN architecture combining CNN and LSTM layers
  - Self-attention mechanism for focusing on important traffic patterns
  - Scalable model design for different traffic volumes
  - Model persistence and loading capabilities

- **Analysis & Visualization**
  - Real-time anomaly detection and scoring
  - Visualization of anomaly patterns over time
  - Configurable detection thresholds
  - Detailed packet-level analysis

## Requirements

### System Requirements
- Python 3.8 or higher
- Sufficient RAM for processing large PCAP files (recommended: 8GB+)
- CUDA-capable GPU (optional, for faster training)

### Dependencies
Install the required dependencies:

```
pip install -r lln_model/requirements.txt --break-system-packages
```

Key dependencies include:
- TensorFlow ≥ 2.12.0: Deep learning framework
- Scapy ≥ 2.4.5: Network packet processing
- NumPy ≥ 1.21.0: Numerical computations
- Scikit-learn ≥ 1.0.0: Data preprocessing
- Matplotlib ≥ 3.5.0: Visualization
- Pandas ≥ 1.3.0: Data manipulation

## Project Structure

```
lln_model/
├── model.py           # LNN model architecture definition
├── preprocessing.py   # Data preprocessing and feature extraction
├── train.py          # Model training pipeline
├── inference.py      # Real-time anomaly detection
├── convert_model.py   # Model format conversion utilities
├── cleanup.sh        # Cleanup script for temporary files
├── requirements.txt  # Project dependencies
└── saved_model/      # Pre-trained models and weights
```

## Building and Running

### 1. Environment Setup

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Install dependencies:

   ```bash
   pip3 install -r lln_model/requirements.txt --break-system-packages
   ```

3. Ensure you have a PCAP file for training/testing:
   - Use your own network capture
   - Download from [Wireshark Samples](https://wiki.wireshark.org/samplecaptures)
   - Use the included sample: `dns-remoteshell.pcap`

### 2. Training the Model

#### Command Line Interface

```bash
python3 lln_model/train.py \
    --pcap-file ./lln_model/dns-remoteshell.pcap \
    --epochs 50 \
    --batch-size 32 \
    --validation-split 0.2 \
    --learning-rate 0.001
```

#### Python API

```python
from lln_model.train import train_model

model, preprocessor = train_model(
    pcap_file="dns-remoteshell.pcap",
    epochs=50,
    batch_size=32,
    validation_split=0.2,
    learning_rate=0.001
)
```

Training parameters:

- `epochs`: Number of training iterations (default: 50)
- `batch_size`: Samples per training batch (default: 32)
- `validation_split`: Fraction of data for validation (default: 0.2)
- `learning_rate`: Model learning rate (default: 0.001)

### 3. Running Anomaly Detection

#### Using Conda Environment

A list of the commands which worked using conda:

```
  conda create -n tf2 python=3.9 tensorflow=2.9 numpy scikit-learn matplotlib joblib scapy pandas h5py -c conda-forge
  conda activate tf2 && python lln_model/train.py --pcap-file ./lln_model/dns-remoteshell.pcap --epochs 50 --batch-size 32 --validation-split 0.2 --learning-rate 0.001
  python -c "import tensorflow as tf; print(tf.__version__)"
  python lln_model/inference.py --pcap-file ./lln_model/dns-remoteshell.pcap --threshold 0.5 
```

```bash
# Activate the conda environment
conda activate tf2

# Run inference with default parameters

python lln_model/inference.py \
    --pcap-file ./lln_model/dns-remoteshell.pcap \
    --threshold 0.5 \
    --output-file anomalies.csv \
    --plot-results

# Or with custom parameters
python lln_model/inference.py \
    --pcap-file ./path/to/your/pcap \
    --threshold 0.7 \
    --output-file custom_anomalies.csv \
    --plot-results
```

#### Command Line Interface
```bash
python3 -m lln_model.inference \
    --pcap-file dns-remoteshell.pcap \
    --threshold 0.5 \
    --output-file anomalies.csv \
    --plot-results
```

#### Python API
```python
from lln_model.inference import AnomalyDetector

detector = AnomalyDetector(
    model_path="lln_model/saved_model",
    threshold=0.5
)

results = detector.detect_anomalies(
    pcap_file="path_to_your_pcap_file.pcap",
    output_file="anomalies.csv"
)

# Visualize results
detector.plot_anomalies(results['anomaly_scores'])
```

Inference parameters:
- `threshold`: Anomaly detection threshold (0.0 to 1.0)
- `output_file`: Path to save detection results
- `plot_results`: Generate visualization of anomalies

## Model Architecture Details

The LNN model implements a hybrid architecture:

1. **Input Layer**
   - Accepts normalized network packet features
   - Supports variable sequence lengths

2. **Feature Extraction**
   - Multiple 1D convolutional layers
   - Max pooling for dimensionality reduction
   - Batch normalization for training stability

3. **Temporal Processing**
   - Bidirectional LSTM layers
   - Self-attention mechanism
   - Dropout for regularization

4. **Classification**
   - Dense layers with ReLU activation
   - Final sigmoid layer for anomaly scoring

## Data Preprocessing Pipeline

1. **Packet Processing**
   - PCAP file parsing using Scapy
   - Feature extraction from packet headers
   - Protocol-specific feature engineering

2. **Data Normalization**
   - Standard scaling of numerical features
   - Categorical feature encoding
   - Sequence padding and truncation

3. **Dataset Preparation**
   - Sliding window sequence generation
   - Train/validation split
   - Batch generation for training

## Best Practices and Tips

1. **Training Data Quality**
   - Use diverse network traffic patterns
   - Include both normal and anomalous samples
   - Clean and preprocess data thoroughly

2. **Model Tuning**
   - Adjust `threshold` based on false positive tolerance
   - Increase `epochs` for better convergence
   - Modify `batch_size` based on available memory

3. **Performance Optimization**
   - Use GPU acceleration when available
   - Adjust sequence length for memory constraints
   - Implement batch processing for large PCAP files

4. **Monitoring and Maintenance**
   - Regularly retrain on new traffic patterns
   - Monitor false positive/negative rates
   - Update detection thresholds as needed

## Troubleshooting

Common issues and solutions:

1. **Memory Errors**
   - Reduce batch size
   - Process PCAP files in chunks
   - Use memory-efficient data types

2. **Training Issues**
   - Check input data quality
   - Adjust learning rate
   - Increase/decrease model complexity

3. **Detection Accuracy**
   - Fine-tune detection threshold
   - Retrain on more relevant data
   - Adjust feature engineering

## Notes

- The model performs best when trained on domain-specific traffic
- Regular retraining helps adapt to new traffic patterns
- Consider using ensemble methods for critical applications
- Monitor system resources during large-scale processing 