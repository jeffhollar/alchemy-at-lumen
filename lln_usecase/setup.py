from setuptools import setup, find_packages

setup(
    name="lln_model",
    version="0.1.0",
    packages=find_packages(),
    install_requires=[
        "tensorflow>=2.12.0",
        "numpy>=1.21.0",
        "scikit-learn>=1.0.0",
        "matplotlib>=3.5.0",
        "joblib>=1.1.0",
        "scapy>=2.4.5",
        "pandas>=1.3.0",
        "h5py>=3.1.0"
    ],
) 