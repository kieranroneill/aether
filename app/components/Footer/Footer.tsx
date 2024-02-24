import { HStack, Link, Text, VStack } from '@chakra-ui/react';
import React, { FC } from 'react';

// constants
import {
  DEFAULT_GAP,
  KIERAN_ONEILL_WEBSITE_LINK,
  LICENSE_LINK,
} from '@app/constants';

// theme
import { theme } from '@app/theme';

const Footer: FC = () => {
  // hooks
  const defaultTextColor: string = theme.colors.whiteAlpha['800'];

  return (
    <VStack
      alignItems="center"
      bg={theme.colors.altBackground}
      px={DEFAULT_GAP * 2}
      py={DEFAULT_GAP}
      spacing={DEFAULT_GAP / 3}
      w="full"
    >
      <HStack
        alignItems="center"
        justifyContent="center"
        spacing={DEFAULT_GAP / 3}
        w="full"
      >
        {/*developed by logo*/}
        <Text color={defaultTextColor}>
          Developed with ❤️ by{' '}
          <Link
            color="primary.50"
            href={KIERAN_ONEILL_WEBSITE_LINK}
            isExternal={true}
          >
            Kieran O&apos;Neill
          </Link>
          {''}.
        </Text>

        {/*license*/}
        <Text color={defaultTextColor}>
          Licensed under{' '}
          <Link color="primary.50" href={LICENSE_LINK} isExternal={true}>
            GNU GPLv3
          </Link>
          {''}.
        </Text>
      </HStack>
    </VStack>
  );
};

export default Footer;
