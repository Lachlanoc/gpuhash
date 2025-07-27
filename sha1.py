import hashlib
import itertools
import string
import time

def brute_force_sha1(target_hash, known_prefix, char_set, max_length=8):
    """
    Brute force a SHA1 hash by trying different alphanumeric combinations
    """
    target_hash = target_hash.lower()
    start_time = time.time()
    attempts = 0

    print(f"Starting brute force with prefix: {known_prefix}")
    print(f"Character set size: {len(char_set)} ({char_set})")
    print(f"Trying combinations up to {max_length} characters")

    # Try combinations of increasing length
    for length in range(1, max_length + 1):
        print(f"Trying {length} character combinations...")

        # Generate all possible combinations of the current length
        for combo in itertools.product(char_set, repeat=length):
            suffix = ''.join(combo)
            potential_flag = f"{known_prefix}{suffix}"

            # Calculate SHA1 hash of the potential flag
            sha1_hash = hashlib.sha1(potential_flag.encode()).hexdigest()

            attempts += 1

            # Check if the hash matches the target
            if sha1_hash == target_hash:
                elapsed = time.time() - start_time
                print(f"Completed in {elapsed:.2f} seconds after {attempts:,} attempts")
                return potential_flag

            # Print progress every million attempts
            if attempts > 0 and attempts % 1000000 == 0:
                elapsed = time.time() - start_time
                rate = attempts / elapsed if elapsed > 0 else 0
                print(f"Tried {attempts:,} combinations... ({rate:.2f} attempts/second)")

    elapsed = time.time() - start_time
    print(f"Search completed in {elapsed:.2f} seconds after {attempts:,} attempts")
    return None

def main():
    # Get inputs from user
    target_hash = input("Enter the SHA1 hash: ")
    known_prefix = input("Enter the known prefix: ")

    # Character set options
    print("\nSelect character set for brute force:")
    print("1. Numbers only (0-9)")
    print("2. Lowercase letters only (a-z)")
    print("3. Uppercase letters only (A-Z)")
    print("4. Numbers and lowercase letters")
    print("5. Numbers and uppercase letters")
    print("6. All letters (a-z, A-Z)")
    print("7. All alphanumeric characters (0-9, a-z, A-Z)")

    choice = input("Enter your choice (1-7): ")

    # Define character set based on user choice
    if choice == '1':
        char_set = string.digits
    elif choice == '2':
        char_set = string.ascii_lowercase
    elif choice == '3':
        char_set = string.ascii_uppercase
    elif choice == '4':
        char_set = string.digits + string.ascii_lowercase
    elif choice == '5':
        char_set = string.digits + string.ascii_uppercase
    elif choice == '6':
        char_set = string.ascii_letters
    elif choice == '7':
        char_set = string.digits + string.ascii_letters
    else:
        print("Invalid choice. Using all alphanumeric characters.")
        char_set = string.digits + string.ascii_letters

    try:
        max_length = int(input("Enter maximum length of characters to try (default: 8): ") or "8")
    except ValueError:
        print("Invalid input. Using default value of 8 characters.")
        max_length = 8

    # Try to find the flag
    flag = brute_force_sha1(target_hash, known_prefix, char_set, max_length)

    if flag:
        print(f"Flag found: {flag}")
    else:
        print("Flag not found. Try increasing max_length or check your inputs.")

if __name__ == "__main__":
    main()

