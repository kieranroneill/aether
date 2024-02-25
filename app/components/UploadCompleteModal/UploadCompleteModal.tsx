import {
  Button,
  Code,
  Heading,
  HStack,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Skeleton,
  Text,
  VStack,
} from '@chakra-ui/react';
import React, { FC } from 'react';
import { IoDownloadOutline } from 'react-icons/io5';

// components
import CopyIconButton from '@app/components/CopyIconButton';

// constants
import { DEFAULT_GAP } from '@app/constants';

// hooks
import useDefaultTextColor from '@app/hooks/useDefaultTextColor';
import useLogger from '@app/hooks/useLogger';
import usePrimaryColorScheme from '@app/hooks/usePrimaryColorScheme';

// types
import type { ILogger } from '@app/types';
import type { IProps } from './types';

// utils
import downloadJSONFile from '@app/utils/downloadJSONFile';

const UploadCompleteModal: FC<IProps> = ({ onClose, uploadResponse }) => {
  // hooks
  const defaultTextColor: string = useDefaultTextColor();
  const logger: ILogger = useLogger();
  const primaryColorScheme: string = usePrimaryColorScheme();
  // handlers
  const handleOnClose = () => onClose();
  const handleDownloadMerkleRoot = () => {
    const _functionName: string = 'handleDownloadMerkleRoot';

    if (!uploadResponse) {
      logger.debug(
        `${UploadCompleteModal.displayName}#${_functionName}: no merkle tree root hash found, ignoring`
      );

      return;
    }

    // create a data uri from the json and download it - use the merkle root as the file name
    downloadJSONFile(uploadResponse.root, uploadResponse);
  };

  return (
    <Modal
      closeOnOverlayClick={false}
      isOpen={!!uploadResponse}
      onClose={onClose}
    >
      <ModalOverlay />

      <ModalContent>
        <ModalHeader>
          <Heading color={defaultTextColor} size="md" textAlign="center">
            {`Files Successfully Uploaded!`}
          </Heading>
        </ModalHeader>

        <ModalBody>
          <VStack spacing={DEFAULT_GAP} w="full">
            {/*description*/}
            <Text color={defaultTextColor} textAlign="left" w="full">
              {`Here is the Merkle Tree root hash that you can use to verify the integrity of a file.`}
            </Text>

            {/*merkle tree root hash*/}
            <HStack alignItems="center" spacing={1} w="full">
              {!uploadResponse ? (
                <Skeleton height="20px" w="full" />
              ) : (
                <>
                  {/*merkle tree root hash*/}
                  <Code borderRadius="md" flexGrow={1} wordBreak="break-word">
                    {uploadResponse.root}
                  </Code>

                  {/*copy button*/}
                  <CopyIconButton
                    ariaLabel={`Copy hash`}
                    value={uploadResponse.root}
                  />
                </>
              )}
            </HStack>
          </VStack>
        </ModalBody>

        <ModalFooter>
          <HStack justifyContent="flex-end" spacing={2} w="full">
            <Button
              colorScheme={primaryColorScheme}
              onClick={handleOnClose}
              variant="outline"
            >
              {`Close`}
            </Button>

            <Button
              colorScheme={primaryColorScheme}
              onClick={handleDownloadMerkleRoot}
              rightIcon={<IoDownloadOutline />}
              variant="solid"
            >
              {`Download`}
            </Button>
          </HStack>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

UploadCompleteModal.displayName = 'UploadCompleteModal';

export default UploadCompleteModal;
