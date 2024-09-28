import type { Meta, StoryObj } from '@storybook/react';

import Stack from '.';

const meta = {
  component: Stack,
  tags: ['autodocs'],
} satisfies Meta<typeof Stack>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Primary: Story = {};