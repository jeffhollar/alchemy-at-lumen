import tensorflow as tf
from tensorflow.keras import layers, Sequential

def create_model(input_shape, num_classes=2, learning_rate=0.001):
    # Ensure input_shape is a tuple
    if isinstance(input_shape, list):
        input_shape = tuple(input_shape)
    
    # Ensure input_shape has the correct format (timesteps, features)
    if len(input_shape) != 2:
        raise ValueError(f"Input shape must be (timesteps, features), got {input_shape}")
    
    # Create a simpler model architecture
    model = Sequential([
        # Input layer with explicit shape and dtype
        layers.Input(shape=input_shape, dtype=tf.float32),
        
        # Simple LSTM layer
        layers.LSTM(64),
        
        # Dense layers
        layers.Dense(32, activation='relu'),
        layers.Dropout(0.3),
        layers.Dense(num_classes, activation='softmax')
    ])
    
    # Compile the model with explicit dtype policy
    tf.keras.mixed_precision.set_global_policy('float32')
    
    model.compile(
        optimizer=tf.keras.optimizers.Adam(learning_rate=learning_rate),
        loss='sparse_categorical_crossentropy',
        metrics=['accuracy']
    )
    return model 