// types
import type { IOptions } from './types';

/**
 * Utility function to truncate a string with an ellipsis in the middle.
 * @param {string} input - the string to truncate.
 * @param {IOptions} options - [optional] options to customise.
 * @returns {string} the truncated string.
 */
export default function truncateString(
  input: string,
  options?: IOptions
): string {
  const defaultLength: number = 5;
  const start: number =
    options && options.start ? options.start : defaultLength;
  const end: number = options && options.end ? options.end : defaultLength;

  return `${input.slice(0, start)}...${input.slice(-end)}`;
}
