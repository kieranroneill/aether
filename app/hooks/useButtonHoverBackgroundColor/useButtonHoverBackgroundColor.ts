import { useColorModeValue } from '@chakra-ui/react';

export default function useButtonHoverBackgroundColor(): string {
  return useColorModeValue('gray.100', 'whiteAlpha.100');
}
