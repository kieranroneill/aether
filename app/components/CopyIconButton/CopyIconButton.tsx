import { Icon, IconButton, Tooltip, useClipboard } from '@chakra-ui/react';
import React, { FC } from 'react';
import { IoCheckmarkOutline, IoCopyOutline } from 'react-icons/io5';

// hooks
import useButtonHoverBackgroundColor from '@app/hooks/useButtonHoverBackgroundColor';
import useDefaultTextColor from '@app/hooks/useDefaultTextColor';

// types
import { IProps } from './types';

const CopyIconButton: FC<IProps> = ({
  ariaLabel,
  tooltipLabel,
  size = 'sm',
  value,
}: IProps) => {
  const { hasCopied, onCopy } = useClipboard(value);
  // hooks
  const buttonHoverBackgroundColor: string = useButtonHoverBackgroundColor();
  const defaultTextColor: string = useDefaultTextColor();
  // handlers
  const handleCopyClick = () => onCopy();

  return (
    <Tooltip
      arrowSize={15}
      hasArrow={true}
      label={tooltipLabel || ariaLabel}
      placement="bottom"
    >
      <IconButton
        _hover={{ bg: buttonHoverBackgroundColor }}
        aria-label={ariaLabel}
        icon={
          hasCopied ? (
            <Icon as={IoCheckmarkOutline} color="green.400" />
          ) : (
            <Icon as={IoCopyOutline} color={defaultTextColor} />
          )
        }
        onClick={handleCopyClick}
        size={size}
        variant="ghost"
      />
    </Tooltip>
  );
};

export default CopyIconButton;
