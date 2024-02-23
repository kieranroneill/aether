import { ChakraProvider } from '@chakra-ui/react';
import { FC } from 'react';

// types
import type { IProps } from './types';

const Providers: FC<IProps> = ({ children, theme }) => {
  return <ChakraProvider theme={theme}>{children}</ChakraProvider>;
};

export default Providers;
