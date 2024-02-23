import { ReactNode } from 'react';

interface ILayoutProps {
  children: ReactNode;
  types: ReactNode;
  params?: Record<string, unknown>;
}

export default ILayoutProps;
