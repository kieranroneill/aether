import {
  ColorModeContextType,
  HStack,
  IconButton,
  Spacer,
  Tooltip,
  useColorMode,
} from '@chakra-ui/react';
import React, { FC } from 'react';
import { IoMoonOutline, IoSunnyOutline } from 'react-icons/io5';

// constants
import { DEFAULT_GAP } from '@app/constants';

// hooks
import useButtonHoverBackgroundColor from '@app/hooks/useButtonHoverBackgroundColor';
import useDefaultTextColor from '@app/hooks/useDefaultTextColor';

const Header: FC = () => {
  const { colorMode, toggleColorMode }: ColorModeContextType = useColorMode();
  // hooks
  const buttonHoverBackgroundColor: string = useButtonHoverBackgroundColor();
  const defaultTextColor: string = useDefaultTextColor();
  // handlers
  const handlerColorChangeClick = () => toggleColorMode();

  return (
    <HStack
      alignItems="center"
      justifyContent="space-between"
      p={DEFAULT_GAP / 2}
      w="full"
    >
      <Spacer />

      {/*color mode toggle*/}
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
