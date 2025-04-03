import onnxruntime as ort
import numpy as np


class ONNXRunner:
    def __init__(self, model_path: str):
        try:
            self.ort_sess = ort.InferenceSession(model_path)
            self.input_name = self.ort_sess.get_inputs()[0].name
        except Exception as e:
            raise RuntimeError(f"Failed to loan ONNX model: {e}")

    def predict(self, features: list[float]) -> float:
        try:
            input_array = np.array([features], dtype=np.float32)
            output = self.ort_sess.run(None, {self.input_name: input_array})
            return float(output[0][0][0])
        except Exception as e:
            raise RuntimeError(f"Inference failed: {e}")
