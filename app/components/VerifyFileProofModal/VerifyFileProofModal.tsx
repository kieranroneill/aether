import {
  Button,
  Code,
  Heading,
  HStack,
  Icon,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Skeleton,
  Spacer,
  Text,
  Textarea,
  VStack,
} from '@chakra-ui/react';
import React, { ChangeEvent, FC, useState } from 'react';
import {
  IoCheckmarkCircleOutline,
  IoCloseCircleOutline,
} from 'react-icons/io5';

// constants
import { DEFAULT_GAP } from '@app/constants';

// hooks
import useDefaultTextColor from '@app/hooks/useDefaultTextColor';
import useLogger from '@app/hooks/useLogger';
import usePrimaryColorScheme from '@app/hooks/usePrimaryColorScheme';
import useSubTextColor from '@app/hooks/useSubTextColor';

// types
import type { ILogger } from '@app/types';
import type { IProps } from './types';

// utils

const VerifyFileProofModal: FC<IProps> = ({ file, onClose }) => {
  // hooks
  const defaultTextColor: string = useDefaultTextColor();
  const logger: ILogger = useLogger();
  const primaryColorScheme: string = usePrimaryColorScheme();
  const subTextColor: string = useSubTextColor();
  // state
  const [rootValue, setRootValue] = useState<string | null>(null);
  const [verified, setVerified] = useState<boolean>(false);
  // handlers
  const handleOnClose = () => onClose();
  const handleRootValueChange = (event: ChangeEvent<HTMLTextAreaElement>) =>
    setRootValue(event.target.value.length > 0 ? event.target.value : null);

  return (
    <Modal closeOnOverlayClick={false} isOpen={!!file} onClose={onClose}>
      <ModalOverlay />

      <ModalContent>
        <ModalHeader>
          <Heading color={defaultTextColor} size="md" textAlign="center">
            {`Verify File Proof`}
          </Heading>
        </ModalHeader>

        <ModalBody>
          <VStack spacing={DEFAULT_GAP} w="full">
            {/*description*/}
            <Text color={defaultTextColor} textAlign="left" w="full">
              {`Here is the Merkle proof for the file.`}
            </Text>

            {/*file name*/}
            <HStack
              alignItems="center"
              justifyContent="space-between"
              spacing={2}
              w="full"
            >
              <Text color={defaultTextColor} fontSize="sm" textAlign="left">
                {`Name:`}
              </Text>

              <Spacer />

              <Text color={subTextColor} fontSize="sm" textAlign="right">
                {file ? file.name : '-'}
              </Text>
            </HStack>

            {/*hash*/}
            <HStack
              alignItems="center"
              justifyContent="space-between"
              spacing={2}
              w="full"
            >
              <Text color={defaultTextColor} fontSize="sm" textAlign="left">
                {`Hash:`}
              </Text>

              <Spacer />

              <Code
                fontSize="sm"
                textAlign="left"
                w="full"
                wordBreak="break-word"
              >
                {file ? file.hash : '-'}
              </Code>
            </HStack>

            {/*proof*/}
            <Code fontSize="sm" textAlign="left" w="full">
              {file ? JSON.stringify(file.proof) : '-'}
            </Code>

            <Text color={defaultTextColor} textAlign="left" w="full">
              {`Enter the Merkle Root and see of the proof is verifiable:`}
            </Text>

            {/*root value input*/}
            <Textarea
              onChange={handleRootValueChange}
              value={rootValue || ''}
            />

            {/*verified status*/}
            <HStack
              alignItems="center"
              justifyContent="center"
              minH={50}
              spacing={2}
              w="full"
            >
              {rootValue ? (
                verified ? (
                  <>
                    <Icon
                      as={IoCheckmarkCircleOutline}
                      color="green.500"
                      h={6}
                      w={6}
                    />

                    <Text
                      color="green.500"
                      textAlign="left"
                    >{`Proof Valid`}</Text>
                  </>
                ) : (
                  <>
                    <Icon
                      as={IoCloseCircleOutline}
                      color="red.500"
                      h={6}
                      w={6}
                    />

                    <Text
                      color="red.500"
                      textAlign="left"
                    >{`Proof Invalid`}</Text>
                  </>
                )
              ) : null}
            </HStack>
          </VStack>
        </ModalBody>

        <ModalFooter>
          <HStack justifyContent="flex-end" spacing={2} w="full">
            <Button
              colorScheme={primaryColorScheme}
              onClick={handleOnClose}
              variant="solid"
            >
              {`Cancel`}
            </Button>
          </HStack>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

VerifyFileProofModal.displayName = 'VerifyFileProofModal';

export default VerifyFileProofModal;
