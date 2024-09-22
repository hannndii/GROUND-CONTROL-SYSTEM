# Menghitung KPK (Kelipatan Persekutuan Terkecil)
def lcm(a, b):
    while b:
        a, b = b, a % b
    return a * b // gcd(a, b)

# Menghitung jumlah bilangan yang habis dibagi oleh divisors
def count_divisible(n, divisors):
    sum = 0
    for i in range(1, 1 << len(divisors)):
        count = bin(i).count("1")
        lcm_value = 1
        for j in range(len(divisors)):
            if i & (1 << j):
                lcm_value = lcm(lcm_value, divisors[j])
        if count % 2 == 1:
            sum += n // lcm_value
        else:
            sum -= n // lcm_value
    return sum

# Daftar pembagi
divisors = [6, 9, 15, 20]

# Menghitung jumlah bilangan yang habis dibagi oleh 6, 9, 15, atau 20
result = count_divisible(1000000000000, divisors)
print(result)