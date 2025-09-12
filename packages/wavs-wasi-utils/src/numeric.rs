/// Macro to convert between Rust's native u128 and WIT tuple<u64, u64> representation.
///
/// Usage:
///   impl_u128_conversions!(my_bindings::exports::my_interface::U128);
///
/// This will implement From traits for bidirectional conversion between
/// the WIT-generated tuple type and Rust's native u128.
#[macro_export]
macro_rules! impl_u128_conversions {
    ($wit_type:ty) => {
        impl From<u128> for $wit_type {
            fn from(value: u128) -> Self {
                let low = value as u64; // Rust guarantees this is truncating to the lower 64 bits
                let high = (value >> 64) as u64;
                Self { value: (high, low) }
            }
        }

        impl From<$wit_type> for u128 {
            fn from(wrapper: $wit_type) -> Self {
                let (high, low) = wrapper.value;
                ((high as u128) << 64) | (low as u128)
            }
        }
    };
}

#[cfg(test)]
mod tests {
    // In theory, something like this would be better, but wit-bindgen only sees the remote type, not the local one.
    // which defeats the purpose of a local test
    //
    // and trying to generate bindings from the types-only crate:
    // 1. doesn't work because it has no world definition
    // 2. just punts the problem since we'd want to write unit tests for any of our wit types
    //
    // wit_bindgen::generate!({
    //     world: "wavs-world",
    //     path: "../../wit-definitions/operator/wit",
    //     //async: true,
    // });
    //
    // so, instead, for this case it's easy to see what the generated type would look like and just define it here

    #[derive(Debug, PartialEq, Copy, Clone)]
    struct FakeU128 {
        value: (u64, u64),
    }

    impl_u128_conversions!(FakeU128);

    #[test]
    fn test_u128_conversion_max() {
        let original: u128 = u128::MAX;
        let wit_repr: FakeU128 = original.into();
        assert_eq!(
            wit_repr,
            FakeU128 {
                value: (u64::MAX, u64::MAX)
            }
        );
        let converted_back: u128 = wit_repr.into();
        assert_eq!(original, converted_back);
    }

    #[test]
    fn test_u128_conversion_endianness() {
        let original = FakeU128 { value: (1, 2) };
        let converted: u128 = original.into();
        let converted_back = FakeU128::from(converted);
        assert_eq!(original, converted_back);
    }
}
