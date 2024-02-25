// types
import type { IUploadResponse } from '@app/types';

interface IProps {
  onClose: () => void;
  uploadResponse: IUploadResponse | null;
}

export default IProps;
