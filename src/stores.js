import { writable } from 'svelte/store';

export const currentStepIndex = writable(0);
export const theme = writable('dark');
export const currentPath = writable('/');

export let previousStepIndices = {prev: 0, current: 0};

currentStepIndex.subscribe(newIndex => {
  previousStepIndices = {
    prev: previousStepIndices.current,
    current: newIndex
  };
});
