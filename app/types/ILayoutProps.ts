import { ReactNode } from 'react';

interface ILayoutProps<Params = Record<string, unknown>> {
  children: ReactNode;
  types: ReactNode;
  params?: Params;
}

export default ILayoutProps;
