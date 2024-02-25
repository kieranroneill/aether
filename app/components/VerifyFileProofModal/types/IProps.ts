// types
import type { IFileResponse } from '@app/types';

interface IProps {
  file: IFileResponse | null;
  onClose: () => void;
}

export default IProps;
