from repository.onnx_runner import ONNXRunner


class LoanApprovalService:
    def __init__(self, predictor: ONNXRunner, threshold: float = 0.5):
        self.predictor = predictor
        self.threshold = threshold

    def evaluate(
        self, annual_income: int, loan_amount: int, loan_term: int, credit_score: int
    ) -> dict:
        features = [annual_income, loan_amount, loan_term, credit_score]
        prediction = self.predictor.predict(features)

        status = "approved" if prediction >= self.threshold else "rejected"
        confidence = abs(prediction - 0.5) * 2

        return {"status": status, "confidence": round(confidence, 4)}
