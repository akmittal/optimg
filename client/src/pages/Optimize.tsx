import { Button, Form, Input, Radio, Select, Slider } from 'antd';
import Title from 'antd/lib/typography/Title';
import { MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';
import React, { ReactElement, useReducer, useState } from 'react';
import './optimize.css';
import Checkbox from 'antd/lib/checkbox/Checkbox';
import { useMutation } from 'react-query';
const { Option } = Select;

interface Transformation {
  format: number;
  quality: number;
}

interface Operation {
  sourcePath: string;
  targetPath: string;
  transformations: Transformation[];
}

interface Props {}

const formItemLayout = {
  labelCol: {
    xs: { span: 24 },
    sm: { span: 4 },
  },
  wrapperCol: {
    xs: { span: 24, display: 'flex' },
    sm: { span: 20, display: 'flex' },
  },
};
const formItemLayoutWithOutLabel = {
  wrapperCol: {
    xs: { span: 24, offset: 0 },
    sm: { span: 20, offset: 0 },
  },
};

function Optimize({}: Props): ReactElement {
  const [state, dispatch] = useReducer(
    function reducer(state: any, action: any) {
      switch (action.type) {
        case 'setSourcePath':
          return { ...state, sourcePath: action.data };

        case 'setTargetPath':
          return { ...state, targetPath: action.data };

        case 'changeQuality':
          return {
            ...state,
            transformations: state.transformations.map(
              (item: any, index: number) => {
                if (index === action.index) {
                  return { ...item, quality: action.data };
                }
                return item;
              }
            ),
          };
        case 'changeFormat':
          return {
            ...state,
            transformations: state.transformations.map(
              (item: any, index: number) => {
                if (index === action.index) {
                  return { ...item, format: action.data };
                }
                return item;
              }
            ),
          };
        case 'addVarient':
          return {
            ...state,
            transformations: [
              ...state.transformations,
              { quality: 75, format: 'avif' },
            ],
          };
        case 'removeVarient':
          return {
            ...state,
            transformations: state.transformations.filter(
              (transformation: any, index: number) => index !== action.index
            ),
          };
      }
    },
    {
      sourcePath: '',
      targetPath: '',
      transformations: [],
    }
  );
  const { isError, isLoading, mutateAsync, isSuccess, error } = useMutation<
    any,
    any
  >(() => {
    let body = JSON.parse(JSON.stringify(state));
    body.transformations = body.transformations.map((transform: any) => ({
      ...transform,
      format: parseInt(transform.format),
    }));
    return fetch('/api/optimize', {
      body: JSON.stringify(body),
      method: 'POST',
    }).then(res => res.json());
  });

  const [form] = Form.useForm();
  return (
    <>
      {isError && error ? <div>An error occurred: {error?.message}</div> : null}

      {isSuccess ? <div>Todo added!</div> : null}
      <div>
        <Title level={3}>Optimize</Title>
        <Form
          layout="vertical"
          form={form}
          initialValues={{ layout: 'vertical' }}
          onValuesChange={() => {}}
          onSubmitCapture={() => {
            console.log(state);
          }}
        >
          <Form.Item label="Source Path">
            <Input
              placeholder="Source Path"
              onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                dispatch({ type: 'setSourcePath', data: e.currentTarget.value })
              }
            />
          </Form.Item>
          <Form.Item label="Destination Path">
            <Input
              placeholder="Destination Path"
              onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                dispatch({ type: 'setTargetPath', data: e.currentTarget.value })
              }
            />
          </Form.Item>
          <Checkbox onChange={() => {}}>Copy Unknown</Checkbox>
          <Checkbox onChange={() => {}}>Monitor</Checkbox>

          {state.transformations.map((field: any, index: number) => (
            <Form.Item
              style={{ display: 'flex' }}
              {...(index === 0 ? formItemLayout : formItemLayoutWithOutLabel)}
              label={index === 0 && <Title level={3}>Varients</Title>}
              required={false}
              key={field.key}
            >
              <Form.Item label="Format" style={{ display: 'inline-block' }}>
                <Select
                  value={field['format']}
                  style={{ width: '160px' }}
                  placeholder="Select Format"
                  onChange={(e: any) =>
                    dispatch({ type: 'changeFormat', index, data: e })
                  }
                >
                  <Option value="10">AVIF</Option>
                  <Option value="2">WebP</Option>
                  <Option value="1">JPEG</Option>
                  <Option value="0">Source Format</Option>
                </Select>
              </Form.Item>
              <Form.Item
                label="Quality"
                style={{
                  display: 'inline-block',
                  margin: '0px 20px',
                  width: 'calc( 50% - 100px) ',
                }}
              >
                <Slider
                  tooltipVisible={true}
                  tooltipPlacement="top"
                  onChange={(e: any) => {
                    console.log(e);
                    dispatch({ type: 'changeQuality', index, data: e });
                  }}
                  value={field['quality']}
                />
              </Form.Item>

              {state.transformations.length > 1 && (
                <MinusCircleOutlined
                  style={{ display: 'inline-block', margin: '35px auto' }}
                  className="dynamic-delete-button"
                  onClick={() => dispatch({ type: 'removeVarient', index })}
                />
              )}
            </Form.Item>
          ))}
          <Form.Item>
            <Button
              type="dashed"
              onClick={() => dispatch({ type: 'addVarient' })}
              style={{ width: '60%' }}
              icon={<PlusOutlined />}
            >
              Add Varient
            </Button>
          </Form.Item>

          <Form.Item>
            <Button
              type="primary"
              onClick={() => {
                mutateAsync();
              }}
            >
              Submit
            </Button>
          </Form.Item>
        </Form>
      </div>
    </>
  );
}

export default Optimize;
