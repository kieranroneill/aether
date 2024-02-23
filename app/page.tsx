'use client';
import { Heading, Text, VStack } from '@chakra-ui/react';
import { NextPage } from 'next';
import React from 'react';

// components
import AetherIcon from '@app/components/AetherIcon';

// constants
import { DEFAULT_GAP } from '@app/constants';

// hooks
import useDefaultTextColor from '@app/hooks/useDefaultTextColor';

const Page: NextPage = () => {
  // hooks
  const defaultTextColor: string = useDefaultTextColor();

  return (
    <VStack
      alignItems="center"
      justifyContent="flex-start"
      flexGrow={1}
      spacing={DEFAULT_GAP}
      w="full"
    >
      {/*icon*/}
      <AetherIcon h={20} w={20} />

      {/*heading*/}
      <Heading color={defaultTextColor} size="lg" textAlign="center" w="full">
        {`Welcome To Aether!`}
      </Heading>

      {/*description*/}
      <Text color={defaultTextColor} size="md" textAlign="center" w="full">
        {process.env.NEXT_PUBLIC_DESCRIPTION}
      </Text>
    </VStack>
  );
};

export default Page;
