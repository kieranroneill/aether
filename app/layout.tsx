import React, { FC } from 'react';

// types
import type { ILayoutProps } from '@app/types';

const RootLayout: FC<ILayoutProps> = ({ children }) => {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
};

export default RootLayout;
