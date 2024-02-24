import {
  ColorModeContextType,
  HStack,
  IconButton,
  Spacer,
  Tooltip,
  useColorMode,
} from '@chakra-ui/react';
import React, { FC } from 'react';
import { IoMenuOutline, IoMoonOutline, IoSunnyOutline } from 'react-icons/io5';

// constants
import { DEFAULT_GAP } from '@app/constants';

// hooks
import useButtonHoverBackgroundColor from '@app/hooks/useButtonHoverBackgroundColor';
import useDefaultTextColor from '@app/hooks/useDefaultTextColor';

// types
import type { IProps } from './types';

const Header: FC<IProps> = ({ onNavigationClick }) => {
  const { colorMode, toggleColorMode }: ColorModeContextType = useColorMode();
  // hooks
  const buttonHoverBackgroundColor: string = useButtonHoverBackgroundColor();
  const defaultTextColor: string = useDefaultTextColor();
  // handlers
  const handlerColorChangeClick = () => toggleColorMode();
  const handleNavigationClick = () => onNavigationClick();

  return (
    <HStack
      alignItems="center"
      justifyContent="space-between"
      p={DEFAULT_GAP / 2}
      w="full"
    >
      {/*open navigation menu button*/}
      <Tooltip label={`Open navigation menu`}>
        <IconButton
          _hover={{ bg: buttonHoverBackgroundColor }}
          aria-label="Open navigation drawer"
          color={defaultTextColor}
          icon={<IoMenuOutline />}
          onClick={handleNavigationClick}
          size="lg"
          variant="ghost"
        />
      </Tooltip>

      <Spacer />

      {/*color mode toggle button*/}
      <Tooltip
        label={`Switch to ${colorMode === 'dark' ? 'light' : 'dark'} mode`}
      >
        <IconButton
          _hover={{ bg: buttonHoverBackgroundColor }}
          aria-label="Change color mode"
          color={defaultTextColor}
          icon={colorMode === 'dark' ? <IoSunnyOutline /> : <IoMoonOutline />}
          onClick={handlerColorChangeClick}
          size="lg"
          variant="ghost"
        />
      </Tooltip>
    </HStack>
  );
};

export default Header;
