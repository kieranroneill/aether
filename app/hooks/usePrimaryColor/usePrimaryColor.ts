import { useColorModeValue } from '@chakra-ui/react';

export default function usePrimaryColor(): string {
  return useColorModeValue('primary.500', 'primary.200');
}
