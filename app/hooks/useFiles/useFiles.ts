import axios, { AxiosResponse } from 'axios';
import { useEffect, useState } from 'react';

// constants
import { FILES_PATH } from '@app/constants';

// hooks
import useLogger from '@app/hooks/useLogger';

// types
import type { IFileResponse, ILogger } from '@app/types';
import type { IUseFilesState } from './types';

export default function useFiles(): IUseFilesState {
  const _functionName: string = 'useFiles';
  // hooks
  const logger: ILogger = useLogger();
  // states
  const [error, setError] = useState<string | null>(null);
  const [files, setFiles] = useState<Record<string, IFileResponse[]> | null>(
    null
  );
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    (async () => {
      let response: AxiosResponse<Record<string, IFileResponse[]>>;

      try {
        response = await axios.get(
          `${process.env.NEXT_PUBLIC_CORE_URL}/${FILES_PATH}`
        );

        setFiles(response.data);
      } catch (error) {
        logger.error(`${_functionName}:`, error);

        setError(error.message);
      }

      setLoading(false);
    })();
  }, []);

  return {
    error,
    files,
    loading,
  };
}
