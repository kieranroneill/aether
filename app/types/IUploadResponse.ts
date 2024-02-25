// types
import IFileDirectoryItem from './IFileDirectoryItem';

interface IUploadResponse {
  directory: IFileDirectoryItem[];
  root: string;
}

export default IUploadResponse;
