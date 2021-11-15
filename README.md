# slp-rep-cat

Hippocampal-Cortical Model Simulation of Category Learning Consolidation over one night of sleep

This is a hippocampal-cortical model of sleep consolidation. The architecture includes C-HORSE, our model of the hippocampus (see [Zhou et. al. 2021](https://www.biorxiv.org/content/10.1101/2021.07.29.454337v1) and [Schapiro et. al. 2017a](https://cb17cd36-5a57-45de-9d66-0b98a3dc5be9.filesusr.com/ugd/b37d16_5edf4f04f8fb4f8eb717d38e4ca42c3e.pdf)) and a neocortical layer as the target of consolidation.

During sleep, the model autonomously replays stimuli it learned  while awake and uses an oscillation-based learning rule during sleep in order to improve its own performance. This version is set up to simulate the structured satellite learning task from [Schapiro et al. 2017b](https://www.nature.com/articles/s41598-017-12884-5.pdf).

## How to run:
1. Clone this repository to your computer. See instructons on how to do this [here](https://docs.github.com/en/free-pro-team@latest/github/creating-cloning-and-archiving-repositories/cloning-a-repository).  
2. ```cd``` into the repository and run the command ```go build  && slp-rep-cat```. When you first run this command, go should automatically download all the packages required to run the model onto your computer. Once the command runs through, a GUI window should open up where you can interact with the model/explore its architecture.  
3. Click the "Train" button on the top left. When you do this, the following will happen:
  i) The model will learn the satellite task in its awake state. You should see blocks of training and testing occuring sequentially until the model hits a learning criterion of 0.66 (66% accuracy on the task).  
  ii) The model will now switch to sleep and will begin replaying the information it just learned during the wake state. At various points the model will fall into periods of high stability which the model will reinforce by contrasting it with immediately following periods of lower stability. The stability measure is displayed at the bottom of the screen as "AvgLaySim", periods of high stability are the "plus phase" of the model and subsequent periods of low stability are the "minus phase".  
  iii) After 30,000 cycles of sleep, the model will switch back to a wake state and will immediately run a test block to measure if there has been an improvement in performance through the learning that occured during sleep.

## Variables that control sleep behaviour:
The model relies on two mechanisms during sleep - (i) Short-term Synaptic Depression which destabilizes item attractors and (ii) Oscillating Inhibition which reveals useful contrastive learning states in destabilized item attractors.
Synaptic depression is controlled by the "inc" and "dec" parameters which specify the rate of increase and recovery from synaptic depression over time, respectively (see line 493).  
Layers in the network recieve either high or low amplitude oscillating inhibition (see line 1128). The amplitude for each is controlled via a sinusoidal equation which can be edited to change the various properties of the oscillations (see line 1556).
