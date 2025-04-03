import torch
from torch.utils.data import DataLoader
from src.dataset import DataFrameDataset


def evaluate(model, X_test, y_test, batch_size):
    model.eval()
    loader = DataLoader(
        DataFrameDataset(X_test, y_test), batch_size=batch_size, shuffle=True
    )
    correct = 0
    with torch.no_grad():
        for X, y in loader:
            pred = model(X)
            correct += ((pred > 0.5) == y).sum().item()
    return correct / len(y_test)
