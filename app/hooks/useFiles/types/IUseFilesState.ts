// types
import { IFileResponse } from '@app/types';

interface IUseFilesState {
  error: string | null;
  files: Record<string, IFileResponse[]> | null;
  loading: boolean;
}

export default IUseFilesState;
