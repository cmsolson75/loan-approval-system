def credit_score_validator(score: int) -> int:
    if score > 900 or score < 300:
        raise ValueError(f"Credit score must be between 300 and 900. Got: {score}")
    return score


def loan_term_validator(term: int) -> int:
    if term < 2 or term > 30:
        raise ValueError("Loan term must be between 2 and 30 years")
    return term


def is_positive(value: int) -> int:
    if value < 0:
        raise ValueError(f"Can't be negative")
    return value
