from torch.utils.data import Dataset
import torch


class DataFrameDataset(Dataset):
    def __init__(self, X, y):
        self.X = X.values.astype("float32")
        self.y = y.values.astype("float32")

    def __len__(self):
        return len(self.y)

    def __getitem__(self, idx):
        return torch.tensor(self.X[idx]), torch.tensor([self.y[idx]])
