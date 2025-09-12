/// Convert a Rust u128 to WIT u128 representation (tuple<u64, u64>)
pub fn u128_to_wit(value: u128) -> (u64, u64) {
    let low = value as u64;
    let high = (value >> 64) as u64;
    (low, high)
}

/// Convert a WIT u128 representation (tuple<u64, u64>) to Rust u128
pub fn wit_to_u128(value: (u64, u64)) -> u128 {
    let (low, high) = value;
    ((high as u128) << 64) | (low as u128)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_u128_conversion() {
        let original: u128 = u128::MAX;
        let wit_repr = u128_to_wit(original);
        assert_eq!(wit_repr, (u64::MAX, u64::MAX));
        let converted_back = wit_to_u128(wit_repr);
        assert_eq!(original, converted_back);
    }
}
