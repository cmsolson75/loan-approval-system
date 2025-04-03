import yaml
from sklearn.model_selection import train_test_split
from src.utils import set_seed
from src.data_utils import load_and_preprocess
from src.train import train_model
from src.eval import evaluate
from src.export import export_onnx

if __name__ == "__main__":
    config = yaml.safe_load(open("config/config.yaml"))
    set_seed(config["training"]["seed"])
    X, y = load_and_preprocess(config["data"]["path"])
    X_train, X_test, y_train, y_test = train_test_split(
        X, y, test_size=config["training"]["test_size"]
    )

    model = train_model(X_train, y_train, config)
    acc = evaluate(model, X_test, y_test, config["training"]["batch_size"])
    print(f"Test Accuracy: {acc * 100:.2f}")

    export_onnx(model, config["model"])
