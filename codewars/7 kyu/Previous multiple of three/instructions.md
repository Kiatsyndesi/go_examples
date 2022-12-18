Given a positive integer n: 0 < n < 1e6, remove the last digit until you're left with a number that is a multiple of three.

Return n if the input is already a multiple of three, and if no such number exists, return null, a similar empty value, or -1.

1. `1      => null`
2. `25     => null`
3. `36     => 36`
4. `1244   => 12`
5. `952406 => 9`