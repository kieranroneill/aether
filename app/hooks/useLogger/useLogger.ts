import { useMemo } from 'react';

// types
import { ILogger, ILogLevel } from '@app/types';

// utils
import createLogger from '@app/utils/createLogger';

export default function useLogger(): ILogger {
  return useMemo<ILogger>(
    () =>
      createLogger((process.env.NEXT_PUBLIC_LOG_LEVEL as ILogLevel) || 'error'),
    []
  );
}
