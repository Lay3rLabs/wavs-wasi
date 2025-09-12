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
                let low = value as u64;
                let high = (value >> 64) as u64;
                (low, high)
            }
        }

        impl From<$wit_type> for u128 {
            fn from(value: $wit_type) -> Self {
                let (low, high) = value;
                ((high as u128) << 64) | (low as u128)
            }
        }
    };
}
