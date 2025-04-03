import torch


def export_onnx(model, config):
    model.eval()
    dummy = torch.randn(1, config["input_dim"])
    torch.onnx.export(
        model,
        dummy,
        config["output_path"],
        input_names=["input"],
        output_names=["output"],
    )
