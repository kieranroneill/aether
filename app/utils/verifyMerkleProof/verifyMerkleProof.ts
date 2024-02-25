// enums
import { LeafPositionEnum } from '@app/enums';

// types
import type { IFileProofItem } from '@app/types';

// utils
import sha256 from '@app/utils/sha256';

/**
 * Verifies that a supplied Merkle root is valid for a Merkle proof.
 * @param {string} root - the root to check against the Merkle proof.
 * @param {IFileProofItem[]} proof - the Merkle proof.
 * @returns {Promise<boolean>} true, if the Merkle root supplied and the calculated Merkle root from the Merkle proof
 * match, false otherwise.
 */
export default async function verifyMerkleProof(
  root: string,
  proof: IFileProofItem[]
): Promise<boolean> {
  let leafHash: string | null = null;
  let value: IFileProofItem;

  // if the merkle tree proof is empty, the root is not in there :P
  if (proof.length <= 0) {
    return false;
  }

  for (let i: number = 0; i < proof.length; i++) {
    value = proof[i];

    // if we are at the first element just use the first hash
    if (i === 0) {
      leafHash = value.hash;

      continue;
    }

    // if the next leaf is a left leaf, create the parent with the right leaf
    if (value.position === LeafPositionEnum.Left) {
      leafHash = await sha256(`${value.hash}${leafHash}`);
      continue;
    }

    // if the next leaf is a right leaf, create the parent with the left leaf
    leafHash = await sha256(`${leafHash}${value.hash}`);
  }

  return leafHash === root;
}
