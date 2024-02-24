'use client';
import { Heading, VStack } from '@chakra-ui/react';
import { NextPage } from 'next';
import React from 'react';

// constants
import { DEFAULT_GAP } from '@app/constants';

// hooks
import useDefaultTextColor from '@app/hooks/useDefaultTextColor';

const VerifyPage: NextPage = () => {
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
      {/*heading*/}
      <Heading color={defaultTextColor} size="lg" textAlign="center" w="full">
        {`Verify`}
      </Heading>
    </VStack>
  );
};

export default VerifyPage;
