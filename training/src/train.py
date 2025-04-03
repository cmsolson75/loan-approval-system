import torch
from torch import nn
from torch.utils.data import DataLoader
from src.dataset import DataFrameDataset
from model.model import LoanPipeline


def train_model(X_train, y_train, config):
    model = LoanPipeline(
        input_dim=config["model"]["input_dim"], output_dim=config["model"]["output_dim"]
    )
    loss_fn = nn.BCELoss()
    optimizer = torch.optim.Adam(model.parameters(), lr=config["training"]["lr"])

    train_loader = DataLoader(
        DataFrameDataset(X_train, y_train),
        batch_size=config["training"]["batch_size"],
        shuffle=True,
    )

    model.train()
    for epoch in range(config["training"]["epochs"]):
        total_loss = 0
        for X, y in train_loader:
            pred = model(X)
            loss = loss_fn(pred, y)
            optimizer.zero_grad()
            loss.backward()
            optimizer.step()
            total_loss += loss.item()
        print(f"Epoch {epoch + 1}: Loss = {total_loss / len(train_loader):.4f}")
    return model
