import torch
import torch.nn as nn
import torch.nn.functional as F


class Scaler(nn.Module):
    def __init__(self):
        super().__init__()
        self.register_buffer(
            "mins",
            torch.tensor(
                [
                    200_000,  # Income Min
                    300_000,  # Amount Min
                    2,  # Term min
                    300,  # Credit Score min
                ],
                dtype=torch.float32,
            ),
        )

        self.register_buffer(
            "maxs",
            torch.tensor(
                [
                    9_900_000,
                    39_500_000,
                    20,
                    900,
                ],
                dtype=torch.float32,
            ),
        )

    def forward(self, x):
        return (x - self.mins) / (self.maxs - self.mins)


class MLP(nn.Module):
    def __init__(self, input_dim: int, output_dim: int):
        super().__init__()
        self.layer1 = nn.Linear(input_dim, 20)
        self.bn1 = nn.BatchNorm1d(20)
        self.hidden = nn.Linear(20, 20)
        self.bn2 = nn.BatchNorm1d(20)
        self.dropout = nn.Dropout(0.25)
        self.out = nn.Linear(20, output_dim)

    def forward(self, x):
        x = self.bn1(F.relu(self.layer1(x)))
        x = self.dropout(self.bn2(F.relu(self.hidden(x))))
        return F.sigmoid(self.out(x))


class LoanPipeline(nn.Module):
    def __init__(self, input_dim: int, output_dim: int):
        super().__init__()
        self.scaler = Scaler()
        self.model = MLP(input_dim, output_dim)

    def forward(self, x):
        x = self.scaler(x)
        return self.model(x)
