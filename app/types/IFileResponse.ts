// types
import IFileDirectoryItem from './IFileDirectoryItem';
import IFileProofItem from './IFileProofItem';

interface IFileResponse extends IFileDirectoryItem {
  proof: IFileProofItem[];
}

export default IFileResponse;
