import torch
import numpy as np
import random
import pandas as pd


def load_and_preprocess(path):
    data = pd.read_csv(path)
    data.columns = data.columns.str.replace(" ", "")
    data["status"] = data.loc[:, "loan_status"].map({" Approved": 1, " Rejected": 0})
    features = ["income_annum", "loan_amount", "loan_term", "cibil_score"]
    return data[features], data["status"]


def set_seed(seed=42):
    torch.manual_seed(seed)
    torch.cuda.manual_seed(seed)
    np.random.seed(seed)
    random.seed(seed)
    torch.backends.cudnn.deterministic = True
    torch.backends.cudnn.benchmark = False
