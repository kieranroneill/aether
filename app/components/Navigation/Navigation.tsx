import {
  Divider,
  Drawer,
  DrawerBody,
  DrawerContent,
  DrawerFooter,
  DrawerHeader,
  DrawerOverlay,
  HStack,
  IconButton,
  Spacer,
  Text,
  Tooltip,
} from '@chakra-ui/react';
import { useRouter } from 'next/navigation';
import React, { FC } from 'react';
import {
  IoCheckmarkCircleOutline,
  IoChevronBackOutline,
  IoCloudUploadOutline,
  IoListOutline,
} from 'react-icons/io5';

// components
import AetherIcon from '@app/components/AetherIcon';
import NavigationLinkItem from './NavigationLinkItem';

// constants
import {
  DEFAULT_GAP,
  FILES_ROUTE,
  INDEX_ROUTE,
  UPLOAD_ROUTE,
  VERIFY_ROUTE,
} from '@app/constants';

// hooks
import useButtonHoverBackgroundColor from '@app/hooks/useButtonHoverBackgroundColor';
import useDefaultTextColor from '@app/hooks/useDefaultTextColor';
import usePrimaryColor from '@app/hooks/usePrimaryColor';

// types
import type { INavigationMenuItem, IProps } from './types';

const Navigation: FC<IProps> = ({ isOpen, onClose }) => {
  const router = useRouter();
  // hooks
  const buttonHoverBackgroundColor: string = useButtonHoverBackgroundColor();
  const defaultTextColor: string = useDefaultTextColor();
  const primaryColor: string = usePrimaryColor();
  // misc
  const navigationMenuItems: INavigationMenuItem[] = [
    {
      icon: IoCloudUploadOutline,
      label: 'Upload',
      route: UPLOAD_ROUTE,
    },
    {
      icon: IoListOutline,
      label: 'Files',
      route: FILES_ROUTE,
    },
    {
      icon: IoCheckmarkCircleOutline,
      label: 'Verify',
      route: VERIFY_ROUTE,
    },
  ];
  // handlers
  const handleOnClose = () => onClose();
  const handleHomeClick = () => {
    router.push(INDEX_ROUTE);
    onClose();
  };

  return (
    <Drawer isOpen={isOpen} placement="left" onClose={onClose}>
      <DrawerOverlay />

      <DrawerContent>
        {/*header*/}
        <DrawerHeader p={0}>
          <HStack
            alignItems="center"
            justifyContent="space-between"
            p={DEFAULT_GAP / 2}
            w="full"
          >
            {/*icon*/}
            <IconButton
              _hover={{ bg: 'transparent' }}
              aria-label="Go To Home"
              color={primaryColor}
              icon={<AetherIcon h={12} w={12} />}
              onClick={handleHomeClick}
              size="2xl"
              variant="ghost"
            />

            <Spacer />

            {/*close navigation menu button*/}
            <Tooltip label={`Close Menu`}>
              <IconButton
                _hover={{ bg: buttonHoverBackgroundColor }}
                aria-label="Close navigation drawer"
                color={defaultTextColor}
                icon={<IoChevronBackOutline />}
                onClick={handleOnClose}
                size="lg"
                variant="ghost"
              />
            </Tooltip>
          </HStack>

          <Divider w="full" />
        </DrawerHeader>

        {/*body*/}
        <DrawerBody p={0}>
          {navigationMenuItems.map(({ icon, label, route }, index) => (
            <NavigationLinkItem
              icon={icon}
              key={`navigation-menu-item-${index}`}
              label={label}
              onClick={handleOnClose}
              route={route}
            />
          ))}
        </DrawerBody>

        {/*footer*/}
        <DrawerFooter>
          <Text
            color={defaultTextColor}
            fontSize="sm"
            textAlign="center"
            w="full"
          >{`v${process.env.NEXT_PUBLIC_VERSION}`}</Text>
        </DrawerFooter>
      </DrawerContent>
    </Drawer>
  );
};

export default Navigation;
