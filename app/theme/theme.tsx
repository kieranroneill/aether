import { extendTheme } from '@chakra-ui/react';
import type { Dict } from '@chakra-ui/utils';

const theme: Dict = extendTheme({
  colors: {
    altBackground: '#303846',
    primary: {
      50: '#8ca9ce',
      100: '#7598c3',
      200: '#5f88b9',
      300: '#4977af',
      400: '#3267a4',
      500: '#185699',
      600: '#134a84',
      700: '#0f3e70',
      800: '#0a325d',
      900: '#07274a',
    },
  },
  config: {
    initialColorMode: 'light',
    useSystemColorMode: true,
  },
  fonts: {
    heading: 'var(--font-lato)',
    body: 'var(--font-lato)',
  },
});

export default theme;
