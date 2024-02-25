/**
 * Convenience function that simple creates an anchor element with a download data URI of the JSON and calls its click
 * handler, then removes it from the DOM.
 * @param {string} fileName - the name of the file to download.
 * @param {unknown} data - the actual JSON to encode in the data URI.
 */
export default function downloadJSONFile<Data = Record<string, unknown>>(
  fileName: string,
  data: Data
): void {
  const dataURI: string = `data:text/json;charset=utf-8,${encodeURIComponent(JSON.stringify(data))}`;
  const anchorElement: HTMLAnchorElement = document.createElement('a');

  anchorElement.setAttribute('href', dataURI);
  anchorElement.setAttribute('download', `${fileName}.json`);

  // append to the element to the body
  document.body.appendChild(anchorElement);

  // download it and remove it
  anchorElement.click();
  anchorElement.remove();
}
