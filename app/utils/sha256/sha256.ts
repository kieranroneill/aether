/**
 * Convenience function that hashes a string using the SHA-256 algorithm and returns as a hexadecimal encoded string.
 * @param {string} input - string to hash.
 * @returns {string} the input hashed using SHA-256 and encoded in hexadecimal.
 */
export default async function sha256(input: string): Promise<string> {
  const hashBuffer: ArrayBuffer = await crypto.subtle.digest(
    'SHA-256',
    new TextEncoder().encode(input)
  );

  return Array.from(new Uint8Array(hashBuffer)) // convert buffer to bytes
    .map((b) => b.toString(16).padStart(2, '0'))
    .join(''); // convert bytes to hex string
}
