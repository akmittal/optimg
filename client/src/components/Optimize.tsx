/** @jsx jsx */
import React, { ReactElement, Fragment, useReducer } from "react";
import { TextField } from "@rmwc/textfield";
import { FormField } from "@rmwc/formfield";
import { Typography } from "@rmwc/typography";
import { Select } from "@rmwc/select";
import { Slider } from "@rmwc/slider";
import { Tooltip } from "@rmwc/tooltip";
import { Button } from "@rmwc/button";
import { Icon } from "@rmwc/icon";
import { css, jsx } from "@emotion/core";
import { IconButton } from "rmwc";
import Header from "./Header";
import { useMutation } from "react-query";

interface Transformation {
  operations: Operation[];
  sourcePath: string;
  targetPath: string;
}
interface Operation {
  format: string;
  scale: number | undefined;
  quality: number;
  height: number | undefined;
  width: number | undefined;
}

const formField = css`
  display: flex;
  flex-direction: column;
  align-items: stretch;
  margin: 10px 15px;
  & > * {
  }
`;

interface Props {}

const optimalTransformations = [
  {
    format: "10",
    scale: undefined,
    quality: 80,
    height: undefined,
    width: undefined,
  },
  {
    format: "2",
    scale: undefined,
    quality: 75,
    height: undefined,
    width: undefined,
  },
];

const initialState: Transformation = {
  sourcePath: "",
  targetPath: "",
  operations: [],
};
type Action =
  | { type: "changeSource"; value: string }
  | { type: "changeTarget"; value: string }
  | { type: "changeFormat"; index: number; format: string }
  | { type: "changeWidth"; index: number; width: number }
  | { type: "changeHeight"; index: number; height: number }
  | { type: "changeQuality"; index: number; quality: number }
  | { type: "changeScale"; index: number; scale: number }
  | { type: "loadOptimal" }
  | { type: "addTransformation" };

function reducer(state: Transformation, action: Action): Transformation {
  switch (action.type) {
    case "changeSource":
      return { ...state, sourcePath: action.value };
    case "changeTarget":
      return { ...state, targetPath: action.value };
    case "changeFormat":
      return {
        ...state,
        operations: [...state.operations].map((opr, index) => {
          if (index === action.index) {
            return { ...opr, format: action.format };
          }
          return opr;
        }),
      };
    case "changeWidth":
      return {
        ...state,
        operations: [...state.operations].map((opr, index) => {
          if (index === action.index) {
            return { ...opr, width: action.width };
          }
          return opr;
        }),
      };
    case "changeHeight":
      return {
        ...state,
        operations: [...state.operations].map((opr, index) => {
          if (index === action.index) {
            return { ...opr, height: action.height };
          }
          return opr;
        }),
      };
    case "changeQuality":
      return {
        ...state,
        operations: [...state.operations].map((opr, index) => {
          if (index === action.index) {
            return { ...opr, quality: action.quality };
          }
          return opr;
        }),
      };
    case "changeScale":
      return {
        ...state,
        operations: [...state.operations].map((opr, index) => {
          if (index === action.index) {
            return { ...opr, scale: action.scale };
          }
          return opr;
        }),
      };
    case "loadOptimal":
      return { ...state, operations: optimalTransformations };
    case "addTransformation":
      return {
        ...state,
        operations: [
          ...state.operations,
          {
            format: "10",
            quality: 80,
            scale: 100,
            width: undefined,
            height: undefined,
          },
        ],
      };
    default:
      throw new Error();
  }
}

function Optimize({}: Props): ReactElement {
  const [state, dispatch] = useReducer(reducer, initialState);
  const [mutate] = useMutation((event: any) => {
    const body = {
      ...state,
      operations: state.operations.map((operation) => ({
        ...operation,
        format: parseInt(operation.format),
      })),
    };
    return fetch("/optimize", { method: "POST", body: JSON.stringify(body) });
  });

  return (
    <div className="flex flex-col ">
      <Header text="Optimize"></Header>
      <FormField css={formField}>
        <label htmlFor="path">Source Path</label>
        <TextField
          name="sourcePath"
          onChange={(e: any) =>
            dispatch({
              type: "changeSource",
              value: e.target.value,
            })
          }
        />
      </FormField>
      <FormField css={formField}>
        <label htmlFor="path">Target Path</label>
        <TextField
          name="targetPath"
          onChange={(e: any) =>
            dispatch({
              type: "changeTarget",
              value: e.target.value,
            })
          }
        />
      </FormField>
      <section className="flex flex-row justify-between mx-4 mt-12 mb-2">
        <Typography use="headline5" className="text-left ">
          {" "}
          Transformations
        </Typography>

        <Button
          label="Add Transformation"
          onClick={(e: any) =>
            dispatch({
              type: "addTransformation",
            })
          }
          icon={<Icon icon={{ icon: "add", size: "large" }} />}
        />
      </section>

      {state.operations.map((operation, index) => (
        <Fragment key={Math.random()}>
          <section className="flex flex-col m-4  rounded-lg shadow-xl">
            <Typography use="headline6">Varient #{index + 1}</Typography>
            <hr />
            <section
              className="flex flex-row flex-wrap"
              css={css`
                & > * {
                  width: 40%;
                }
              `}
            >
              <FormField css={formField}>
                <label htmlFor="path">Format: </label>
                <Select
                  name={index.toString()}
                  options={{
                    null: "Keep Original",
                    10: "AVIF",
                    2: "WEBP",
                    1: "JPEG",
                    3: "PNG",
                  }}
                  onChange={(e: any) =>
                    dispatch({
                      type: "changeFormat",
                      index: index,
                      format: e.target.value,
                    })
                  }
                  value={operation.format}
                  defaultValue="10"
                />
              </FormField>
              <FormField css={formField}>
                <label htmlFor="path">Quality: </label>
                <Slider
                  value={operation.quality}
                  name={index.toString()}
                  min={1}
                  onChange={(e: any) =>
                    dispatch({
                      type: "changeQuality",
                      index: index,
                      quality: e.target.value,
                    })
                  }
                  // onInput={(evt) => console.log(evt.detail.value)}
                  // discrete
                  step={1}
                />
                <div className="-mt-6">{operation.quality}%</div>
              </FormField>
              <FormField css={formField}>
                <label htmlFor="path">Height: </label>
                <label>
                  <TextField
                    name={`${index}-height`}
                    onChange={(e: any) =>
                      dispatch({
                        type: "changeHeight",
                        index: index,
                        height: e.target.value,
                      })
                    }
                    type="number"
                    defaultValue={0}
                    value={operation.height}
                    min={0}
                  />{" "}
                  px
                </label>
              </FormField>
              <FormField css={formField}>
                <label htmlFor="path">
                  Width:
                  <Tooltip content="0 denoted default width" showArrow>
                    <IconButton icon="info" />
                  </Tooltip>
                </label>
                <label>
                  <TextField
                    type="number"
                    name="width"
                    defaultValue={0}
                    value={operation.width}
                    onChange={(e: any) =>
                      dispatch({
                        type: "changeWidth",
                        index: index,
                        width: e.target.value,
                      })
                    }
                    min={0}
                  />{" "}
                  px
                </label>
              </FormField>
              <FormField css={formField}>
                <label htmlFor="path">Scale: </label>
                <label>
                  <TextField
                    type="number"
                    name="scale"
                    value={operation.scale}
                    defaultValue={100}
                    onChange={(e: any) =>
                      dispatch({
                        type: "changeScale",
                        index: index,
                        scale: e.target.value,
                      })
                    }
                    max={10000}
                  />{" "}
                  %
                </label>
              </FormField>
            </section>
          </section>
        </Fragment>
      ))}
      <Button
        outlined
        onClick={() =>
          dispatch({
            type: "loadOptimal",
          })
        }
      >
        Load Optimal Defaults
      </Button>
      <Button raised onClick={mutate}>
        Optimize
      </Button>
    </div>
  );
}

export default Optimize;
