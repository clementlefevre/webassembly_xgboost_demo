# -*- coding: utf-8 -*-
"""xgb_binary_classifier.ipynb

"""


import pandas as pd
import xgboost as xgb
import numpy as np
import collections
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score, confusion_matrix
from sklearn.utils import shuffle


# Set column dtypes for Pandas
COLUMN_NAMES = collections.OrderedDict({
    'var_1': 'category' ,
     'var_2': np.int64,
       'predictor':np.int64
})


# Load data into Pandas
data = pd.read_csv('data.csv',
  index_col=False,
  dtype=COLUMN_NAMES
)
data = data.dropna()
data = shuffle(data, random_state=2)
data.head()

# Label preprocessing
labels = data['predictor'].values

# See the distribution of approved / denied classes (0: denied, 1: approved)
print(data['predictor'].value_counts())

data = data.drop(columns=['predictor'])

# Convert categorical columns to dummy columns
dummy_columns = list(data.dtypes[data.dtypes == 'category'].index)
data = pd.get_dummies(data, columns=dummy_columns)

# Preview the data
data.head()

"""## Train the XGBoost model"""

# Split the data into train / test sets
x,y = data,labels
x_train,x_test,y_train,y_test = train_test_split(x,y)

x_test.to_csv()

# Train the model, this will take a few minutes to run
bst = xgb.XGBClassifier(
    objective='binary:logistic'
)

bst.fit(x_train, y_train)

# Get predictions on the test set and print the accuracy score
y_pred = bst.predict(x_test)
acc = accuracy_score(y_test, y_pred.round())
print(acc, '\n')

# Print a confusion matrix
print('Confusion matrix:')
cm = confusion_matrix(y_test, y_pred.round())
cm = cm / cm.astype(np.float).sum(axis=1)
print(cm)

# Save the model so that we can deploy it
bst.save_model('model.bst')