// import { ChartData, ChartOptions, ChartTypeRegistry } from "chart.js"
import { Dispatch, createContext, useReducer } from "react";
// import _ from "lodash";

// interface AuxProps {
//     children: React.ReactNode
// }
// const initialOptions: ChartOptions<keyof ChartTypeRegistry> =  {
//     responsive: true,
//     maintainAspectRatio: true,
//     indexAxis: "x" as const, scales: {
//         x: { stacked: false, },
//         y: { stacked: false, }
//     },
//     plugins:{
//         legend: { position: "top" as const },
//         title: { 
//             display: true,
//             text: "Chart.js Chart"
//         }
//     }
// }
// interface OptionAction {
//     type : "vertical" | "horizontal" | "stacked" | "custom";
//     modifiedOptions?: ChartOptions<keyof ChartTypeRegistry>
// }
// function optionReducer (
//     options: ChartOptions<keyof ChartTypeRegistry>,
//     action: OptionAction
// ): ChartOptions<keyof ChartTypeRegistry> {
//     switch (action.type) {
//         case "vertical": {
//             const newOptions = _.cloneDeep(options) as ChartOptions<"bar">
//             if (newOptions.scales && newOptions.scales.x && newOptions.scales.x.stacked && newOptions.scales.y && newOptions.scales.y.stacked) {
//                 newOptions.scales.x.stacked = false;
//                 newOptions.scales.y.stacked = false;
//             }
//             if (newOptions.indexAxis = "y") {
//                 newOptions.indexAxis = 'x'
//             }
//             return newOptions;
//         }
//         case "horizontal": {
//             const newOptions = _.cloneDeep(options) as ChartOptions<"bar">
//             if (newOptions.scales && newOptions.scales.x && !newOptions.scales.x.stacked && newOptions.scales.y && !newOptions.scales.y.stacked) {
//                 newOptions.scales.x.stacked = true;
//                 newOptions.scales.y.stacked = true;
//             }
//             return newOptions;
//         }
//         case "stacked": {
//             const newOptions = _.cloneDeep(options) as ChartOptions<"bar">
//             if (newOptions.scales && newOptions.scales.x && !newOptions.scales.x.stacked && newOptions.scales.y && !newOptions.scales.y.stacked) {
//                 newOptions.scales.x.stacked = true;
//                 newOptions.scales.y.stacked = true;
//             }
//             return newOptions;
//         }
//         case "custom": {
//             const newOptions = _.cloneDeep(options) as ChartOptions<"bar">;
//             return {...newOptions, ...action.modifiedOptions};
//         }
//     }
// }
// const OptionsContext = createContext<ChartOptions<keyof ChartTypeRegistry>>(initialOptions)
// const OptionDispachContext = createContext<Dispatch<any> | undefined> (undefined)
// export const OptionsProvider: React.FC<AuxProps> = ({children}) => {
//     const [optionsProv, dispachOptionProv]: [ChartOptions<keyof ChartTypeRegistry>, Dispatch<any>] = useReducer(optionReducer, initialOptions)
//     return (
//         <OptionsContext.Provider value = {optionsProv}>
//             <OptionDispachContext.Provider value={dispachOptionProv}>
//                 {children}
//             </OptionDispachContext.Provider>
//         </OptionsContext.Provider>
//     );
// };
// export enum Months {
//     JANUARY, FEBUARY, MARCH, APRIL, MAY, JUNE, JULY, AUGUST, SEPTEMBER, OCTOBER, NOVEMBER, DECEMBER
// }
// interface DataAction {
//     type: "addDataset" | "removeDataset" | "addMonth" | "removeMonth" | "custom",
//     modifiedData?: ChartData<keyof ChartTypeRegistry>
// }
// const labels = [Months[0]]
// export const initialData: ChartData<keyof ChartTypeRegistry> = {
//     labels: labels,
//     datasets: [
//         {
//             label: "Dataset 1",
//             data: labels.map(() => Math.floor(Math.random() * 100)),
//             backgroundColor: "rgba(255,99,132,0.5)",
//         },
//         {
//             label: "Dataset 2",
//             data: labels.map(() => Math.floor(Math.random() * 100)),
//             backgroundColor: "rgba(54,162,235,0.5)",
//         }
//     ],
// };
// const DataContext = createContext<ChartData<keyof ChartTypeRegistry>>(initialData)
// const DataDispachContext = createContext<Dispatch<any> | undefined>(undefined);

// function dataReducer(
//     data: ChartData<keyof ChartTypeRegistry>,
//     action: DataAction
// ): ChartData<keyof ChartTypeRegistry> {
//     switch(action.type) {
//         case "addDataset": {
//             const newData = _.cloneDeep(data)
//             const newDataset = {
//                 label:  `Dataset ${newData.datasets.length + 1}`,
//                 data: newData && newData.labels ? newData.labels.map(() => Math.floor(Math.random() * 100)),
//                 backgroundColor: `rgba(${Math.floor(Math.random() * 255)}, ${Math.floor(Math.random() * 255)}, ${Math.floor(Math.random() * 255)}, 0.5)`
//             }

//         }
//     }
// }