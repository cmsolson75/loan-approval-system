import pandas as pd


def load_and_preprocess(path):
    data = pd.read_csv(path)
    data.columns = data.columns.str.replace(" ", "")
    data["status"] = data.loc[:, "loan_status"].map({" Approved": 1, " Rejected": 0})
    features = ["income_annum", "loan_amount", "loan_term", "cibil_score"]
    return data[features], data["status"]
