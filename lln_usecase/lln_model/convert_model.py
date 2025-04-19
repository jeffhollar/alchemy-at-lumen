import tensorflow as tf
import torch
import numpy as np
from lln_model.model import LongNeuralNetwork

def convert_keras_to_pytorch(keras_model_path, pytorch_model_path):
    # Load the Keras model
    keras_model = tf.keras.models.load_model(keras_model_path)
    
    # Create a PyTorch model with the same architecture
    input_shape = keras_model.input_shape[1:]
    pytorch_model = LongNeuralNetwork(input_shape)
    
    # Convert weights
    for layer in keras_model.layers:
        if isinstance(layer, tf.keras.layers.Dense):
            # Convert dense layer weights
            weights = layer.get_weights()
            if len(weights) > 0:
                pytorch_model.dense1.weight.data = torch.from_numpy(weights[0].T)
                if len(weights) > 1:
                    pytorch_model.dense1.bias.data = torch.from_numpy(weights[1])
    
    # Save the PyTorch model
    torch.save(pytorch_model, pytorch_model_path)
    print(f"Model converted and saved to {pytorch_model_path}")

if __name__ == "__main__":
    keras_path = "lln_model/saved_model/model.keras"
    pytorch_path = "lln_model/saved_model/model.pt"
    convert_keras_to_pytorch(keras_path, pytorch_path) 