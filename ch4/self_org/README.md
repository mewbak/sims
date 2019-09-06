Back to [All Sims](https://github.com/CompCogNeuro/sims) (also for general info and executable downloads)

# Introduction

This model illustrates how self-organizing learning emerges from the interactions between the following factors (as discussed in the [Learning
Chapter](https://grey.colorado.edu/CompCogNeuro/index.php/CCNBook/Learning#Self-Organizing_Learning:_Long_Time_Scales_and_the_BCM_Model)):

* **Inhibitory competition** -- only the most strongly driven neurons get over the inhibitory threshold, and can get active. These are the ones whose current synaptic weights best fit ("detect") the current input pattern.

* **Rich get richer** positive feedback loop -- due to the nature of the learning function, only those neurons that actually get active are capable of learning (when receiver activity y = 0, then xy = 0 too, and the XCAL dWt function is 0 at 0). Thus, the neurons that already detect the current input the best are the ones that get to further strengthen their ability to detect these inputs. This is the essential insight that Hebb had with why the Hebbian learning function should strengthen an "engram".

* **homeostasis** to balance the positive feedback loop -- if left unchecked, the rich-get-richer dynamic ends up with a few units dominating everything, and as a result, all the inputs get categorized into one useless, overly-broad category ("everything"). The homeostatic mechanism in BCM helps fight against this by raising the floating threshold for highly active neurons, causing their weights to decrease, and restoring a balance. Similarly, under-active neurons experience net weight increases that get them participating and competing more effectively.

The net result is the development of a set of neural detectors that
relatively evenly cover the space of different inputs patterns, with
systematic categories that encompass the statistical regularities. For
example, cats like milk, and dogs like bones, and we can learn this just
by observing the reliable co-occurrence of cats with milk and dogs with
bones. This kind of reliable co-ocurrence is what we mean by
"statistical regularity". See [Hebbian
Learning](/CCNBook/Learning/Hebbian "wikilink") for a very simple
illustration of why Hebbian-style learning mechanisms capture patterns
of co-occurrence. It is really just a variant on the basic maxim that
"things that fire together, wire together".

In this exploration, the network learns about a simple world that
consists purely of horizontal and vertical lines, with these lines
appearing always in combination with other lines. The clear objective of
self-organizing learning in this case is to extract the underlying
statistical regularity that these lines exist as reliable collections of
pixels, and it would be much more efficient to encode this world in
terms of the lines, instead of individual pixels.

Let's take a look at the network. The 5x5 input projects to a hidden
layer of 20 units, which are all fully connected to the input with
random initial weights, and have standard FFFB inhibitory competition
dynamics operating amongst them.

Because viewing the pattern of weights over all the hidden units will be
of primary concern as the network learns, we have a special grid view
showing on the upper left of the network display, which displays the
weights for all hidden units. In addition, there is a graph view in the
upper right, which will display key information as the network learns.

Let's see the environment the network will be experiencing.

This will bring up a display showing the first 10 training items on the
left, which are composed of the elemental horizontal and vertical lines
shown in the grid view display on the right. You can use the narrow
vertical violet scrollbar for the left grid view to scroll through all
45 of the patterns (you may have to click the red arrow first to be able
to grab the scrollbar). These 45 input patterns represent all unique
pairwise combinations of vertical and horizontal lines. Thus, there are
no real correlations between the lines, with the only reliable
correlations being between the pixels that make up a particular line. To
put this another way, each line can be thought of as appearing in a
number of different randomly related contexts (i.e., with other lines).

It should be clear that if we computed the correlations between
individual pixels across all of these images, everything would be
equally (weakly) correlated with everything else. Thus, learning must be
conditional on the particular type of line for any meaningful
correlations to be extracted. We will see that this conditionality will
simply self-organize through the interactions of the learning rule and
the FFFB inhibitory competition. Note also that because two lines are
present in every image, the network will require at least two active
hidden units per input, assuming each unit is representing a particular
line.

You should see one of the event patterns containing two lines in the
input of the network, and a pattern of roughly two or more active hidden
units (the FFFB inhibition is very approximate in determining how many
units are active).

After 30 *epochs* (passes through all 45 different events in the
environment) of learning, the network will stop. You should have noticed
that the grid weights view was updated as the training proceeded. This
grid view display shows all of the network's weights. The larger-scale
5x4 grid is topographically arranged in the same layout as the Hidden
layer of the network. Within each of these 20 grid elements is a smaller
5x5 grid representing the input units, showing the weights for each
unit. By clicking on the hidden units in the network window with the
r.wt variable selected, you should be able to verify this
correspondence.

As training proceded, the weights came to more and more clearly reflect
the lines present in the environment. Thus, individual units developed
*selective* representations of the correlations present within
individual lines, or two lines in some cases. The BCM-based XCAL
learning algorithm does not alter weights from inactive inputs, so it
tends to accumulate a bit of "cruft" (a historical trace of the learning
process) in the weights, but the weights to the dominant inputs for each
unit get very strong and stand out. This lack of learning to inactive
inputs (which differs significantly from more standard forms of Hebbian
learning, e.g., as used in the previous version of the Leabra learning
algorithm) is not only biologically supported, but also significantly
increases the overall storage capacity of the network, by reducing
interference from prior learning.

These line representations developed as a result of the interaction
between learning and inhibitory competition as follows. Early on, the
units that won the inhibitory competition were those that happened to
have larger random weights for the input pattern. Learning then tuned
these weights to be more selective for that input pattern, causing them
to be more likely to respond to that pattern and others that overlap
with it (i.e., other patterns sharing one of the two lines). To the
extent that the weights are stronger for one of the two lines in the
input, the unit will be more likely to respond to inputs having this
line, and the weights will continue to increase. If a unit gets over
active, then its long-term average activity level, which sets the
floating threshold for the BCM-style learning, will result in weight
decreases that help to refine its response properties.

The dynamics of the inhibitory competition are critical for the
self-organizing effect, causing different units to specialize on
different lines. Just as in Darwinian evolution, competition drives the
pressure to specialize..

The net result of this self-organizing learning is a *combinatorial*
distributed representation, where each input pattern is represented as
the combination of the two line features present therein. This is the
"obvious" way to represent such inputs, but you should appreciate that
the network nevertheless had to discover this representation through the
somewhat complex self-organizing learning procedure.

Notice that in general two or more units are strongly activated by each
input pattern, with the extra activation reflecting the fact that some
lines are coded by multiple units.

Another thing to notice in the weights shown in the grid view is that
some units are obviously not selective for anything. These "loser" units
(also known as "dead" units) were never reliably activated by any input
feature, and thus did not experience much learning. It is typically
quite important to have such units lying around, because
self-organization requires some "elbow room" during learning to sort out
the allocation of units to stable correlational features. Having more
hidden units also increases the chances of having a large enough range
of initial random selectivities to seed the self-organization process.
The consequence is that you need to have more units than is minimally
necessary, and that you will often end up with leftovers (plus the
redundant units mentioned previously).

From a biological perspective, we know that the cortex does not produce
a lot of new cortical neurons in adults, so we conclude that in general
there is probably an excess of neural capacity relative to the demands
of any given learning context. Thus, it is useful to have these leftover
and redundant units, because they constitute a "reserve" that could
presumably get activated if new features were later presented to the
network (e.g., diagonal lines). We are much more suspicious of
algorithms that require precisely tuned quantities of hidden units to
work properly (more on this later).

# Unique Pattern Statistic

Although looking at the weights is informative, we could use a more
concise measure of how well the network's internal model matches the
underlying structure of the environment. We one such measure is plotted
in the graph view as the network learns.

This log shows the results of the **unique pattern statistic** (shown as
uniq_pats in the graph), which records the number of unique hidden unit
activity patterns that were produced as a result of probing the network
with all 10 different types of horizontal and vertical lines (presented
individually). Thus, there is a separate testing process which, after
each epoch of learning, tests the network on all 10 lines, records the
resulting hidden unit activity patterns (with the FFFB inhibition
cranked up to 5 so that typically 1 unit is active, so we can find the
*most active* response to each input), and then counts up the number of
unique such patterns (subject to thresholding so we only care about
binary patterns of activity).

The logic behind this measure is that if each line is encoded by (at
least) one distinct hidden unit, then this will show up as a unique
pattern. If, however, there are units that encode two or more lines
together (which is not as efficient of a model of this environment),
then this will not result in a unique representation for these lines,
and the resulting measure will be lower. Thus, to the extent that this
statistic is less than 10, the internal model produced by the network
does not fully capture the underlying independence of each line from the
other lines. Note, however, that the unique pattern statistic does not
care if *multiple* hidden units encode the same line (i.e., if there is
redundancy across different hidden units) -- it only cares that the
*same* hidden unit not encode *two different* lines.

Also, for most runs, if you use the lower level of inhibition used
during training, there will always be a unique pattern of activity -- in
the brain, distributed representations (as discussed in the
[Network](/CCNBook/Network "wikilink") chapter) are much more efficient
for encoding unique patterns via different patterns of active units --
so this uniq_pats statistic is really a strict, simple measure of
learning performance.

The performance of the model on any given run can be quite variable,
depending on the random initial weights. Almost always the uniq_pats
statistic is above 5, and often it is a perfect 10, and typically it
climbs up over training. Because of this variability, we need to run
multiple batches of training to get a better sense of how well the
network learns in general.

After the 8 training runs, the batch view shows summary statistics about
the average (mean), maximum, and minimum of the unique pattern statistic
at the end of each network training run. The last column contains a
count of the number of times that a "perfect 10" on the unique pattern
statistic was recorded.

# Parameter Manipulations

Now, let's explore the effects of some of the parameters in the control
panel.

One thing that is a bit unrealistic about this model is the lack of any
activity at all in the units that are off. In the real brain, inactive
neurons always have some low level of activity. This can affect the
extent to which weights decrease to the less active inputs, potentially
leading to cleaner overall patterns of weights.

Another factor that will affect learning are the rate constants
controlling the long term average activation, which then sets the
floating threshold level. These are in the field in the control panel,
particularly (the most important parameter) and . You can experiment
with changing these values, and looking at both the weights and
uniq_pats stat. Reducing the value from 1.5 to 1, which reduces the
long term average values, reduces the homeostatic pressure, and causes
more "hogging" (rich getting richer with no counter-pressure).
Increasing to 2.0 makes the units less likely to represent multiple
lines in one unit, by increasing this homeostatic pressure.

You can also look at the (long term average activity) unit variable
(found in the NetView control panel) in the Network as it trains. With
default parameters, you should observe that the active units all have
roughly the same average activity, while the inactive ones have much
lower averages. You can also experiment with changing the learning rate.

In conclusion, this exercise should give you a feel for the dynamics
that underlie self-organizing learning, and also for the importance of
how the floating threshold level and homeostasis dynamic plays a key
role in this form of learning.

