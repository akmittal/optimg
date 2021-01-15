import 'two-up-element';

import React, { ReactElement } from 'react';
import Title from 'antd/lib/typography/Title';

declare global {
    namespace JSX {
      interface IntrinsicElements {
        'two-up': React.DetailedHTMLProps<React.HTMLAttributes<HTMLElement>, HTMLElement>; // Normal web component; // Normal web component
       
      }
    }
  }

interface Props {}

export default function Compare({}: Props): ReactElement {
  return (
    <div>
        <div className="flex justify-between items-baseline">
            <Title level={4}>Original Image</Title>
            <Title level={4}>Converted Image</Title>
        </div>
      <two-up>
        <div><img src="https://via.placeholder.com/1024/0000FF" className="w-full" /></div>
        <div><img src="https://via.placeholder.com/1024/FF00FF" className="w-full" /></div>
      </two-up>
    </div>
  );
}
