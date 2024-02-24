import { Button, HStack, Icon, Text } from '@chakra-ui/react';
import { useRouter } from 'next/navigation';
import React, { FC } from 'react';
import { IconType } from 'react-icons';
import { IoArrowForwardOutline } from 'react-icons/io5';

// constants
import { DEFAULT_GAP, NAVIGATION_ITEM_HEIGHT } from '@app/constants';

// hooks
import useButtonHoverBackgroundColor from '@app/hooks/useButtonHoverBackgroundColor';
import useDefaultTextColor from '@app/hooks/useDefaultTextColor';

interface IProps {
  icon: IconType;
  label: string;
  onClick: () => void;
  route: string;
}

const NavigationLinkItem: FC<IProps> = ({ icon, label, onClick, route }) => {
  const router = useRouter();
  // hooks
  const buttonHoverBackgroundColor: string = useButtonHoverBackgroundColor();
  const defaultTextColor: string = useDefaultTextColor();
  //misc
  const iconSize: number = 4;
  // handlers
  const handleOnClick = () => {
    router.push(route);
    onClick();
  };

  return (
    <Button
      _hover={{
        bg: buttonHoverBackgroundColor,
      }}
      borderRadius={0}
      fontSize="md"
      h={NAVIGATION_ITEM_HEIGHT}
      justifyContent="start"
      onClick={handleOnClick}
      p={0}
      variant="ghost"
      w="full"
    >
      <HStack
        alignItems="center"
        justifyContent="space-between"
        p={DEFAULT_GAP / 2}
        spacing={DEFAULT_GAP - 2}
        w="full"
      >
        {/*icon*/}
        <Icon as={icon} color={defaultTextColor} h={iconSize} w={iconSize} />

        {/*label*/}
        <Text
          color={defaultTextColor}
          flexGrow={1}
          fontSize="md"
          textAlign="left"
        >
          {label}
        </Text>

        {/*arrow forward icon*/}
        <Icon
          as={IoArrowForwardOutline}
          color={defaultTextColor}
          h={iconSize}
          w={iconSize}
        />
      </HStack>
    </Button>
  );
};

export default NavigationLinkItem;
